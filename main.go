package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"scaffold/config"
	"scaffold/dao/mysql"
	"scaffold/dao/redis"
	"scaffold/logger"
	"scaffold/routes"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
    // load configuration
    config.Init()

    // initialize logger
    logger.Init()
    defer zap.L().Sync()

    // initialize database connection
    mysql.Init()
    defer mysql.Close()

    // initialize Redis connection
    redis.Init()
    defer redis.Close()

    // register routes
    r := routes.Setup()

    // start server (graceful shutdown)
    server := &http.Server{
        Addr: fmt.Sprintf(":%d", viper.GetInt("server.port")),
        Handler: r,
    }

    go func() {
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