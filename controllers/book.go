package controllers

import (
	"fmt"
	"net/http"
	"q3/database"
	"q3/repository"
	"q3/structs"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	var (
		result gin.H
	)

	categories, err := repository.GetAllBooks(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var book structs.Book
	printOut := ""

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.DecideThickness()
	regex, err := regexp.Compile(`/((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)/`)
	if err != nil {
		fmt.Println(err.Error())
	}
	isMatch := regex.MatchString(book.ImageUrl)

	switch {
	case book.ReleaseYear < 1980:
		printOut = "tahun rilis minimal 1980"
		fallthrough
	case !isMatch:
		printOut = printOut + ", URL gambar tidak valid"
	default:
		err = repository.InsertBook(database.DbConnection, book)
		if err != nil {
			panic(err)
		}
		printOut = "Success Insert Book"
	}

	c.JSON(http.StatusOK, gin.H{
		"result": printOut,
	})
}

func UpdateBook(c *gin.Context) {
	var book structs.Book
	printOut := ""
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.DecideThickness()
	regex, err := regexp.Compile(`/((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)/`)
	if err != nil {
		fmt.Println(err.Error())
	}
	isMatch := regex.MatchString(book.ImageUrl)

	book.Id = int(id)

	switch {
	case book.ReleaseYear < 1980:
		printOut = "tahun rilis minimal 1980"
		fallthrough
	case !isMatch:
		if !isMatch {
			printOut = printOut + ", URL gambar tidak valid"
		}
	default:
		err = repository.UpdateBook(database.DbConnection, book)
		if err != nil {
			panic(err)
		}
		printOut = "Success Update Book"
	}

	c.JSON(http.StatusOK, gin.H{
		"result": printOut,
	})

}

func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, err := strconv.Atoi(c.Param("id"))

	book.Id = int(id)

	err = repository.DeleteBook(database.DbConnection, book)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Book",
	})

}
