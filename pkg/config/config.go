package config

import (
	"dayu.com/gindemo/pkg/util"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	Cfg = viper.New()

	AppName string
	RunMode string

	LogPath  string
	LogLevel string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	c := flag.String("c", "./conf/app.toml", "Do not found the conf file")
	flag.Parse()

	configPath := util.AbsPath(*c)
	log.Println("conf path is: ", configPath)

	// load conf
	Cfg.AddConfigPath(".")
	Cfg.SetConfigFile(configPath)
	if err := Cfg.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error conf file: %s \n", err))
	}

	Cfg.WatchConfig()
	Cfg.OnConfigChange(func(in fsnotify.Event) {
		log.Println("loading change conf file: ", in.Name)
	})

	loadApp()
	loadServer()
}

func loadApp() {
	AppName = Cfg.GetString("app.name")
	RunMode = Cfg.GetString("app.mode")
}

func loadServer() {
	HTTPPort = Cfg.GetInt("server.port")
	ReadTimeout = time.Duration(Cfg.GetInt("server.read_timeout")) * time.Second
	WriteTimeout = time.Duration(Cfg.GetInt("server.write_timeout")) * time.Second
}
