package controllers

import (
	"inven-el/database"
	"inven-el/repository"
	"inven-el/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {
	var (
		result gin.H
	)

	item, err := repository.GetAllItems(database.DbConnection)

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

func GetItem(c *gin.Context) {
	var (
		result gin.H
	)
	id := c.Param("id")

	item, err := repository.GetItem(database.DbConnection, id)

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

func InsertItemList(c *gin.Context) {
	var item structs.ItemList

	err := c.ShouldBindJSON(&item)
	if err != nil {
		panic(err)
	}

	err = repository.InsertItemList(database.DbConnection, item)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Item",
	})
}

func UpdateItemList(c *gin.Context) {
	var item structs.ItemList
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&item)
	if err != nil {
		panic(err)
	}

	item.Id = int64(id)

	err = repository.UpdateItemList(database.DbConnection, item)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Item",
	})

}

func DeleteItemList(c *gin.Context) {
	var item structs.ItemList
	id, err := strconv.Atoi(c.Param("id"))

	item.Id = int64(id)

	err = repository.DeleteItemList(database.DbConnection, item)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Item",
	})

}
