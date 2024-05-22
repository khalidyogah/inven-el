package controllers

import (
	"inven-el/database"
	"inven-el/repository"
	"inven-el/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDepartment(c *gin.Context) {
	var (
		result gin.H
	)

	department, err := repository.GetAllDepartment(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": department,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetDepartment(c *gin.Context) {
	var (
		result gin.H
	)
	id := c.Param("id")

	department, err := repository.GetDepartment(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": department,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertDepartment(c *gin.Context) {
	var department structs.Department

	err := c.ShouldBindJSON(&department)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create jwt token",
		})
		return
	}

	defer PanicHandler(c, "Failed insert item")
	err = repository.InsertDepartment(database.DbConnection, department)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Item",
	})
}

func UpdateDepartment(c *gin.Context) {
	var department structs.Department
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&department)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create jwt token",
		})
		return
	}

	department.Id = int64(id)

	err = repository.UpdateDepartment(database.DbConnection, department)

	defer PanicHandler(c, "Failed update item")
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Item",
	})

}

func DeleteDepartment(c *gin.Context) {
	var department structs.Department
	id, err := strconv.Atoi(c.Param("id"))

	defer PanicHandler(c, "Failed delete item")
	department.Id = int64(id)

	err = repository.DeleteDepartment(database.DbConnection, department)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Item",
	})

}
