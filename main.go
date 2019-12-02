package main

import (
	"context"
	"dayu.com/gindemo/app/config"
	"dayu.com/gindemo/router"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	c := flag.String("c", "./conf/app.toml", "Do not found the config file")
	flag.Parse()

	// load config
	viper.AddConfigPath("./conf")
	viper.SetConfigFile(*c)
	log.Println("config path is: ", *c)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("loading change config file: ", in.Name)
	})
}

func main() {
	gin.SetMode(viper.GetString("app.mode"))

	e := gin.New()
	// load router
	router.SetupRouter(e)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("app.port")),
		Handler:      e,
		ReadTimeout:  config.HTTPReadTimeout * time.Second,
		WriteTimeout: config.HTTPWriteTimeout * time.Second,
	}

	log.Println("Port: " + viper.GetString("app.port") + "	Pid: " + fmt.Sprintf("%d", os.Getpid()))

	// listen server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("Http server listen: %s\n", err))
		}
	}()

	// receive sign and exit
	gracefulExitApp(server)
}

// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
func gracefulExitApp(server *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
