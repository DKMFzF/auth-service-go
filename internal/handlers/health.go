package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerHealth(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
