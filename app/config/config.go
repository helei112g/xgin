package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

const (
	// 超时时间
	HTTPReadTimeout  = 120
	HTTPWriteTimeout = 120
)

var Cfg = viper.New()

func InitConfig() {
	c := flag.String("c", "./config/app.toml", "Do not found the config file")
	flag.Parse()

	// load config
	Cfg.AddConfigPath(".")
	Cfg.SetConfigFile(*c)
	log.Println("config path is: ", *c)
	if err := Cfg.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	Cfg.WatchConfig()
	Cfg.OnConfigChange(func(in fsnotify.Event) {
		log.Println("loading change config file: ", in.Name)
	})
}
