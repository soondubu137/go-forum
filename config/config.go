package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf is a global variable that holds the configuration values
var Conf = new(Config)

type Config struct {
    Name    string   `mapstructure:"name"`
    Mode    string   `mapstructure:"mode"`
    Host    string   `mapstructure:"host"`
    Port    int      `mapstructure:"port"`
    Version string   `mapstructure:"version"`
    *LogConfig       `mapstructure:"log"`
    *MySQLConfig     `mapstructure:"mysql"`
    *RedisConfig     `mapstructure:"redis"`
    *SnowflakeConfig `mapstructure:"snowflake"`
}

type LogConfig struct {
    Level      string `mapstructure:"level"`
    MaxSize    int    `mapstructure:"max_size"`
    MaxAge     int    `mapstructure:"max_age"`
    MaxBackups int    `mapstructure:"max_backups"`
    Filename   string `mapstructure:"filename"`
}

type MySQLConfig struct {
    Host         string `mapstructure:"host"`
    Port         int    `mapstructure:"port"`
    User         string `mapstructure:"user"`
    Password     string `mapstructure:"password"`
    DB           string `mapstructure:"db"`
    MaxOpenConns int    `mapstructure:"max_open_conns"`
    MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
    PoolSize int    `mapstructure:"pool_size"`
}

type SnowflakeConfig struct {
    StartTime string `mapstructure:"start_time"`
    MachineID int64  `mapstructure:"machine_id"`
}

func Init() {
    viper.SetConfigFile("config.yaml")
    viper.AddConfigPath(".")

    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }

    if err := viper.Unmarshal(Conf); err != nil {
        panic(err)
    }

    // hot reload any change in the config file
    viper.WatchConfig()
    viper.OnConfigChange(func(in fsnotify.Event) {
        fmt.Println("Config file changed:", in.Name)
        // don't forget to unmarshal the changed configuration
        if err := viper.Unmarshal(Conf); err != nil {
            panic(err)
        }
    })
}