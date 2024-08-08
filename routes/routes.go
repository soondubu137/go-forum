package routes

import (
	"net/http"

	"github.com/SoonDubu923/go-forum/controller"
	"github.com/SoonDubu923/go-forum/logger"
	"github.com/SoonDubu923/go-forum/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
    // if in release mode, use the release mode of Gin
    if mode == gin.ReleaseMode {
        gin.SetMode(gin.ReleaseMode)
    }

    r := gin.New()
    r.Use(logger.GinLogger(), logger.GinRecovery(true))

    v1 := r.Group("/api/v1")

    // register routes here
    v1.POST("/register", controller.RegisterHandler)
    v1.POST("/login", controller.LoginHandler)
    v1.Use(middleware.AuthMiddleware())
    {
        v1.GET("/community", controller.CommunityHandler)
    }

    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound, gin.H{"status": "Error", "message": "Page not found"})
    })
    return r
}