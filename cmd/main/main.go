package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"car-recommendation-service/internal/shared/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	configName := "main"
	if err := config.SetupConfig(configName); err != nil {
		log.Fatalf("failed to load the config file: %s", err.Error())
	}

	router := gin.Default()
	router.LoadHTMLGlob(viper.GetString("pattern"))
	path1 := viper.GetString("path1")
	root1 := viper.GetString("root1")
	path2 := viper.GetString("path2")
	root2 := viper.GetString("root2")
	path3 := viper.GetString("path3")
	root3 := viper.GetString("root3")
	path4 := viper.GetString("path4")
	root4 := viper.GetString("root4")
	router.Static(path1, root1)
	router.Static(path2, root2)
	router.Static(path3, root3)
	router.Static(path4, root4)

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

	router = MakeNewRouter(router)

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
