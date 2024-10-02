package server

import (
	"InfSec/internal/transport/rest/handlers"
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	handlers.InitTemplates(r)

	v1 := r.Group("v1/")
	{
		v1.GET("/ping", handlers.Ping)
	}

	r.GET("", handlers.IndexHandler)
}
