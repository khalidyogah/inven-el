package controllers

import (
	"inven-el/database"
	"inven-el/repository"
	"inven-el/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var department structs.Department

	err := c.ShouldBindJSON(&department)
	if err != nil {
		panic(err)
	}

	err = repository.InsertDepartment(database.DbConnection, department)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Item",
	})
}

func Login(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	err = repository.Login(database.DbConnection, department)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Item",
	})

}
