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
	router.GET("/item-con/:id", controllers.GetItemCondition)

	router.GET("/item-type", controllers.GetAllTypes)
	router.POST("/item-type", controllers.InsertItemType)
	router.PUT("/item-type/:id", controllers.UpdateItemType)
	router.DELETE("/item-type/:id", controllers.DeleteItemType)
	router.GET("/item-type/:id", controllers.GetItemType)

	router.GET("/department", controllers.GetAllDepartment)
	router.POST("/department", controllers.InsertDepartment)
	router.PUT("/department/:id", controllers.UpdateDepartment)
	router.DELETE("/department/:id", controllers.DeleteDepartment)
	router.GET("/department/:id", controllers.GetDepartment)

	return router
}
