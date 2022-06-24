package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}

func ErrRouter(c *gin.Context) {
	c.String(http.StatusBadRequest, "url err")
}
