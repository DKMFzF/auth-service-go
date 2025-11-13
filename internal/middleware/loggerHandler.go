package middleware

import (
	"auth/internal/logs"
	"strconv"

	gin "github.com/gin-gonic/gin"
)

func LoggerHandler(log *logs.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		log.Info(
			"%s", "Request "+
				c.Request.RequestURI+
				" Response Code "+strconv.Itoa(c.Writer.Status()),
		)
	}
}
