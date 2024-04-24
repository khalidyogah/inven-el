package controllers

import (
	"inven-el/database"
	"inven-el/repository"
	"inven-el/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllItemConditions(c *gin.Context) {
	var (
		result gin.H
	)

	item, err := repository.GetAllItemConditions(database.DbConnection)

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

func GetItemCondition(c *gin.Context) {
	var (
		result gin.H
	)
	id := c.Param("id")

	item, err := repository.GetItemCondition(database.DbConnection, id)

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

func InsertItemCondition(c *gin.Context) {
	var item structs.ItemCondition

	err := c.ShouldBindJSON(&item)
	if err != nil {
		panic(err)
	}

	err = repository.InsertItemCondition(database.DbConnection, item)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Item",
	})
}

func UpdateItemCondition(c *gin.Context) {
	var item structs.ItemCondition
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&item)
	if err != nil {
		panic(err)
	}

	item.Id = int64(id)

	err = repository.UpdateItemCondition(database.DbConnection, item)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Item",
	})

}

func DeleteItemCondition(c *gin.Context) {
	var item structs.ItemCondition
	id, err := strconv.Atoi(c.Param("id"))

	item.Id = int64(id)

	err = repository.DeleteItemCondition(database.DbConnection, item)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Item",
	})

}
