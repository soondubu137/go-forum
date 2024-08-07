package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SoonDubu923/go-forum/config"
	"github.com/SoonDubu923/go-forum/dao/mysql"
	_ "github.com/SoonDubu923/go-forum/dao/redis"
	"github.com/SoonDubu923/go-forum/logger"
	"github.com/SoonDubu923/go-forum/pkg/snowflake"
	"github.com/SoonDubu923/go-forum/routes"

	"go.uber.org/zap"
)

func main() {
    // load configuration
    config.Init()

    // initialize logger
    if err := logger.Init(config.Conf.LogConfig); err != nil {
        panic(fmt.Sprintf("init logger failed: %v", err))
    }
    defer zap.L().Sync()

    // initialize database connection
    if err := mysql.Init(config.Conf.MySQLConfig); err != nil {
        zap.L().Fatal("mysql.Init failed", zap.Error(err))
    }
    defer mysql.Close()

    // initialize Redis connection
    // if err := redis.Init(config.Conf.RedisConfig); err != nil {
    //     zap.L().Fatal("redis.Init failed", zap.Error(err))
    // }
    // defer redis.Close()

    // register routes
    r := routes.Setup()

    // initialize snowflake node
    if err := snowflake.Init(config.Conf.SnowflakeConfig); err != nil {
        zap.L().Fatal("snowflake.Init failed", zap.Error(err))
    }

    // start server (graceful shutdown)
    server := &http.Server{
        Addr: fmt.Sprintf("%s:%d", config.Conf.Host, config.Conf.Port),
        Handler: r,
    }

    go func() {
        zap.L().Info("Starting server", zap.String("addr", server.Addr))
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            zap.L().Fatal("listen: %s\n", zap.Error(err))
        }
    }()

    // graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
    <-quit
    zap.L().Info("Shutdown Server ...")
    // allow 30 seconds for existing connections to finish
    ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        zap.L().Fatal("Server Shutdown:", zap.Error(err))
    }
    zap.L().Info("Server exiting")
}
