package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Logger())

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "goods",
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
