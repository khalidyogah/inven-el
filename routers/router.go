package routers

import (
	"q3/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/bangun-datar/:bangun", controllers.HitungBangunDatar)

	router.GET("/categories", controllers.GetAllCategories)
	router.POST("/categories", controllers.InsertCategory)
	router.PUT("/categories/:id", controllers.UpdateCategory)
	router.DELETE("/categories/:id", controllers.DeleteCategory)
	router.GET("/categories/:id/book", controllers.GetCategory)

	router.GET("/books", controllers.GetAllBooks)
	router.POST("/books", controllers.InsertBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}
