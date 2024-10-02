package server

import (
	handlers2 "InfSec/src/transport/rest/handlers"
	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	handlers2.InitTemplates(r)

	v1 := r.Group("v1/")
	{
		v1.GET("/ping", handlers2.Ping)
	}

	r.GET("", handlers2.IndexHandler)
}
