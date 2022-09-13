package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Authorization(c *gin.Context) {
	token := c.GetHeader("TOKEN")
	if token == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}
	if token != os.Getenv("ACCESS_TOKEN") {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}
	c.Next()
}
