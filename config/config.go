package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() {
    viper.SetConfigFile("config.yaml")
    viper.AddConfigPath(".")

    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }

    viper.WatchConfig()
    viper.OnConfigChange(func(in fsnotify.Event) {
        fmt.Println("Config file changed:", in.Name)
    })
}