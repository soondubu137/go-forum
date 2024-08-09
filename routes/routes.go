package routes

import (
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
        v1.GET("/communities", controller.CommunityHandler)
        v1.GET("/community/:id", controller.CommunityDetailHandler)
        v1.POST("/publish", controller.PublishHandler)
        v1.GET("/post/:id", controller.PostDetailHandler)
        v1.GET("/posts", controller.PostListHandler)
    }

    r.NoRoute(func(c *gin.Context) {
        controller.ErrorResponse(c, controller.CodeNotFound)
    })
    return r
}