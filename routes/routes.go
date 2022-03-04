package routes

import (
	"go-rest-api/controller"
	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {
	articlesGroup := r.Group("/api/v1/articles")
	articleController := controller.Articles{}
	articlesGroup.GET("", articleController.FildAll)
	articlesGroup.GET("/:id", articleController.FildOne)
	articlesGroup.POST("",articleController.Create)
}
