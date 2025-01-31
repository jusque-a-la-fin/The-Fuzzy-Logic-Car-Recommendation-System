package main

import (
	"car-recommendation-service/internal/shared/config"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var reg *regexp.Regexp

func main() {
	configName := "selection"
	if err := config.SetupConfig(configName); err != nil {
		log.Fatalf("failed to load the config file: %s", err.Error())
	}

	reg = regexp.MustCompile(viper.GetString("expression"))
	router := gin.Default()

	tmpl := template.New("")
	pattern1 := viper.GetString("pattern1")
	tmpl, err := tmpl.ParseGlob(pattern1)
	if err != nil {
		panic(err)
	}

	pattern2 := viper.GetString("pattern2")
	tmpl, err = tmpl.ParseGlob(pattern2)
	if err != nil {
		panic(err)
	}

	pattern3 := viper.GetString("pattern3")
	tmpl, err = tmpl.ParseFiles(pattern3)
	if err != nil {
		panic(err)
	}

	router.SetHTMLTemplate(tmpl)
	path1 := viper.GetString("path1")
	root1 := viper.GetString("root1")
	path2 := viper.GetString("path2")
	root2 := viper.GetString("root2")
	path3 := viper.GetString("path3")
	root3 := viper.GetString("root3")
	path4 := viper.GetString("path4")
	root4 := viper.GetString("root4")
	path5 := viper.GetString("path5")
	root5 := viper.GetString("root5")
	path6 := viper.GetString("path6")
	root6 := viper.GetString("root6")
	router.Static(path1, root1)
	router.Static(path2, root2)
	router.Static(path3, root3)
	router.Static(path4, root4)
	router.Static(path5, root5)
	router.Static(path6, root6)

	origin1 := viper.GetString("origin1")
	method1 := viper.GetString("method1")
	method2 := viper.GetString("method2")
	method3 := viper.GetString("method3")
	method4 := viper.GetString("method4")
	header1 := viper.GetString("header1")
	header2 := viper.GetString("header2")
	header3 := viper.GetString("header3")
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{origin1},
		AllowMethods: []string{method1, method2, method3, method4},
		AllowHeaders: []string{header1, header2, header3},
	}))

	carsDir := viper.GetString("carRoot")
	id1 := viper.GetString("id1")
	id2 := viper.GetString("id2")
	id3 := viper.GetString("id3")
	deepestDirs := getDeepestDirsperformDFS(carsDir)
	for _, dir := range deepestDirs {
		id1Val, id2Val, id3Val, err := GetIDs(dir)
		if err != nil {
			log.Fatalf("error from GetIDs, package `main`, selection service %v", err)
		}
		path := fmt.Sprintf("%s=%s&%s=%s&%s=%s", id1, id1Val, id2, id2Val, id3, id3Val)
		router.Static(path, dir)
	}

	storageCon, err := grpc.NewClient(
		viper.GetString("addr1"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("can't create a new gRPC `channel`", err)
	}
	defer storageCon.Close()

	selectionCon, err := grpc.NewClient(
		viper.GetString("addr2"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("can't create a new gRPC `channel`", err)
	}
	defer selectionCon.Close()

	router = MakeNewRouter(router, storageCon, selectionCon)
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/main")
	})

	httpServer := &http.Server{
		Addr:           viper.GetString("addr"),
		Handler:        router,
		ReadTimeout:    1000 * time.Second,
		WriteTimeout:   1000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Can't listen or serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

func getDeepestDirsperformDFS(rootDir string) []string {
	var deepestDirs []string
	var maxDepth int
	performDFS(rootDir, 0, &deepestDirs, &maxDepth)
	return deepestDirs
}

func performDFS(dir string, depth int, deepestDirs *[]string, maxDepth *int) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			performDFS(filepath.Join(dir, file.Name()), depth+1, deepestDirs, maxDepth)
		}
	}
	if depth > *maxDepth {
		*maxDepth = depth
		*deepestDirs = []string{dir}
	} else if depth == *maxDepth {
		*deepestDirs = append(*deepestDirs, dir)
	}
}

func GetIDs(root string) (string, string, string, error) {
	matches := reg.FindStringSubmatch(root)
	if len(matches) == 4 {
		id1 := matches[1]
		id2 := matches[2]
		id3 := matches[3]
		return id1, id2, id3, nil
	}
	return "", "", "", fmt.Errorf("no matches found")
}
