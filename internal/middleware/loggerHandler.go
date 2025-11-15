package middleware

import (
	//"auth/internal/logs"
	"fmt"
	"strconv"

	gin "github.com/gin-gonic/gin"
)

func LoggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		fmt.Printf(
			"%s", "Request "+
				c.Request.RequestURI+
				" Response Code "+strconv.Itoa(c.Writer.Status()),
		)
	}
}
