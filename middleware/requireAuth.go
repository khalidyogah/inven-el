package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequiredAuth(c *gin.Context) {
	fmt.Println("mid")
	c.Next()
}
