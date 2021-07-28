package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouteRegister(router *gin.RouterGroup) {
	router.POST("/", func(context *gin.Context) {
		fmt.Println(context.RemoteIP())
	})

	router.POST("/login", func(context *gin.Context) {
		fmt.Println(context.RemoteIP())
	})
}
