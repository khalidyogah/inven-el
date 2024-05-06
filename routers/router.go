package routers

import (
	"inven-el/controllers"
	"inven-el/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/signup/", controllers.CreateAccount)
	router.POST("/login/", controllers.Login)
	// router.GET("/validate/", middleware.RequiredAuth, controllers.Validate)

	router.GET("/item-list/", controllers.GetAllItems)
	router.POST("/item-list", middleware.RequiredAuth, controllers.InsertItemList)
	router.PUT("/item-list/:id", middleware.RequiredAuth, controllers.UpdateItemList)
	router.DELETE("/item-list/:id", middleware.RequiredAuth, controllers.DeleteItemList)
	router.GET("/item-list/:id", controllers.GetItem)

	router.GET("/item-con", controllers.GetAllItemConditions)
	router.POST("/item-con", middleware.RequiredAuth, controllers.InsertItemCondition)
	router.PUT("/item-con/:id", middleware.RequiredAuth, controllers.UpdateItemCondition)
	router.DELETE("/item-con/:id", middleware.RequiredAuth, controllers.DeleteItemCondition)
	router.GET("/item-con/:id", controllers.GetItemCondition)

	router.GET("/item-type", controllers.GetAllTypes)
	router.POST("/item-type", middleware.RequiredAuth, controllers.InsertItemType)
	router.PUT("/item-type/:id", middleware.RequiredAuth, controllers.UpdateItemType)
	router.DELETE("/item-type/:id", middleware.RequiredAuth, controllers.DeleteItemType)
	router.GET("/item-type/:id", controllers.GetItemType)

	router.GET("/department", controllers.GetAllDepartment)
	router.POST("/department", middleware.RequiredAuth, controllers.InsertDepartment)
	router.PUT("/department/:id", middleware.RequiredAuth, controllers.UpdateDepartment)
	router.DELETE("/department/:id", middleware.RequiredAuth, controllers.DeleteDepartment)
	router.GET("/department/:id", controllers.GetDepartment)

	return router
}
