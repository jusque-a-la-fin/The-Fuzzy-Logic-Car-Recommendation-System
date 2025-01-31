package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"car-recommendation-service/internal/shared/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	configName := "search"
	if err := config.SetupConfig(configName); err != nil {
		log.Fatalf("failed to load the config file: %s", err.Error())
	}

	router := gin.Default()
	tmpl := template.New("")
	pattern := viper.GetString("pattern")
	tmpl, err := tmpl.ParseGlob(pattern)
	if err != nil {
		panic(err)
	}

	file := viper.GetString("file")
	tmpl, err = tmpl.ParseFiles(file)
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

	router.Static(path1, root1)
	router.Static(path2, root2)
	router.Static(path3, root3)
	router.Static(path4, root4)
	router.Static(path5, root5)

	origin1 := viper.GetString("origin1")
	method1 := viper.GetString("method1")
	method2 := viper.GetString("method2")
	method3 := viper.GetString("method3")
	method4 := viper.GetString("method4")
	header1 := viper.GetString("header1")
	header2 := viper.GetString("header2")
	header3 := viper.GetString("header3")
	allowCredentialsStr := viper.GetString("allowCredentials")
	allowCredentials, err := strconv.ParseBool(allowCredentialsStr)
	if err != nil {
		log.Fatalf("error from `ParseBool` function, package `strconv`: %v", err)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{origin1},
		AllowMethods:     []string{method1, method2, method3, method4},
		AllowHeaders:     []string{header1, header2, header3},
		AllowCredentials: allowCredentials,
	}))

	storageCon, err := grpc.NewClient(
		viper.GetString("addr1"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("can't create a new gRPC `channel`", err)
	}
	defer storageCon.Close()

	surveyCon, err := grpc.NewClient(
		viper.GetString("addr2"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("can't create a new gRPC `channel`", err)
	}
	defer surveyCon.Close()

	carsCon, err := grpc.NewClient(
		viper.GetString("addr3"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("can't create a new gRPC `channel`", err)
	}
	defer carsCon.Close()

	router = MakeNewRouter(router, storageCon, surveyCon, carsCon)

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
