package main

import (
	"context"
	"dayu.com/gindemo/framework/config"
	"dayu.com/gindemo/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := gin.New()

	// middleware
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	gin.SetMode(config.RunMode)
	// load router
	router.SetupRouter(e)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.HTTPPort),
		Handler:      e,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	log.Println("Port: " + config.Cfg.GetString("app.port") + "	Pid: " + fmt.Sprintf("%d", os.Getpid()))

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
