package routers

import (
	"inven-el/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/item-list/", controllers.GetAllItems)
	router.POST("/item-list", controllers.InsertItemList)
	router.PUT("/item-list/:id", controllers.UpdateItemList)
	router.DELETE("/item-list/:id", controllers.DeleteItemList)
	router.GET("/item-list/:id", controllers.GetItem)

	router.GET("/item-con", controllers.GetAllItemConditions)
	router.POST("/item-con", controllers.InsertItemCondition)
	router.PUT("/item-con/:id", controllers.UpdateItemCondition)
	router.DELETE("/item-con/:id", controllers.DeleteItemCondition)
	router.GET("/item-con/:id/book", controllers.GetItemCondition)

	// router.GET("/books", controllers.GetAllBooks)
	// router.POST("/books", controllers.InsertBook)
	// router.PUT("/books/:id", controllers.UpdateBook)
	// router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}
