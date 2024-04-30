package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gustasousagh/jobs-api-golang/handler"
	"github.com/gustasousagh/jobs-api-golang/middleware"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	middleware.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("opening", handler.ShowOpeningHandler)
		v1.POST("opening", handler.CreateOpeningHandler)
		v1.DELETE("opening", handler.DeleteOpeningHandler)
		v1.PUT("opening", handler.UpdateOpeningHandler)
		v1.GET("openings", handler.ListOpeningHandler)
		v1.POST("register", handler.RegisterHandler)
		v1.POST("login", handler.Login)
		v1.GET("validate", middleware.Auth, handler.Validate)
	}
}
