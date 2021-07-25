package main

import (
	"github.com/gin-gonic/gin"
	"github.com/happiness-daily-tli/happiness-daily-api/common"
	"github.com/happiness-daily-tli/happiness-daily-api/users"
	"gorm.io/gorm"
)

func main() {
	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	r.Use(gin.Logger())

	v1 := r.Group("/api/v1")
	users.UsersRegister(v1.Group("/users"))

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

//func Migrate(db *gorm.DB) {
//	users.AutoMigrate()
//	db.AutoMigrate(&articles.ArticleModel{})
//	db.AutoMigrate(&articles.TagModel{})
//	db.AutoMigrate(&articles.FavoriteModel{})
//	db.AutoMigrate(&articles.ArticleUserModel{})
//	db.AutoMigrate(&articles.CommentModel{})
//}
