package controllers

import (
	"inven-el/database"
	"inven-el/repository"
	"inven-el/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTypes(c *gin.Context) {
	var (
		result gin.H
	)

	item, err := repository.GetAllTypes(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": item,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetItemType(c *gin.Context) {
	var (
		result gin.H
	)
	id := c.Param("id")

	item, err := repository.GetItemType(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": item,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertItemType(c *gin.Context) {
	var item structs.ItemType

	err := c.ShouldBindJSON(&item)
	if err != nil {
		panic(err)
	}

	err = repository.InsertItemType(database.DbConnection, item)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Item",
	})
}

func UpdateItemType(c *gin.Context) {
	var item structs.ItemType
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&item)
	if err != nil {
		panic(err)
	}

	item.Id = int64(id)

	err = repository.UpdateItemType(database.DbConnection, item)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Item",
	})

}

func DeleteItemType(c *gin.Context) {
	var item structs.ItemType
	id, err := strconv.Atoi(c.Param("id"))

	item.Id = int64(id)

	err = repository.DeleteItemType(database.DbConnection, item)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Item",
	})

}
