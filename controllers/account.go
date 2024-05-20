package controllers

import (
	"inven-el/database"
	"inven-el/repository"
	"inven-el/structs"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func PanicHandler(c *gin.Context, message string) {
	if message := recover(); message != nil {
		c.JSON(http.StatusOK, gin.H{
			"result": message,
		})
	}
}

func CreateAccount(c *gin.Context) {
	//get the email pass req body
	var user structs.User

	defer PanicHandler(c, "Failed create account")
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	//create user
	err = repository.CreateAccount(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed create account",
		})
		return
	}

	//respond
	c.JSON(http.StatusOK, gin.H{
		"result": "Success create account",
	})
}

func Login(c *gin.Context) {
	//get email pass from body
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//look up requested user
	//compare sent in pass with saved user pass hash
	err = repository.Login(database.DbConnection, user)

	//generate jwt token (subject, expiration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create jwt token",
		})
		return
	}

	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})

	// c.JSON(http.StatusOK, gin.H{
	// 	"token": tokenString,
	// })

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
