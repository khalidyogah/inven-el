package middleware

import (
	"fmt"
	"inven-el/database"
	"inven-el/repository"
	"inven-el/structs"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequiredAuth(c *gin.Context) {
	//get cookie of req
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode/validate
	//parse
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//dont forget to validate the alg is what u expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find user with token sub
		var user structs.User

		err := repository.FindUser(database.DbConnection, claims["sub"], &user)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// attch to req
		c.Set("user", user)

		//continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
