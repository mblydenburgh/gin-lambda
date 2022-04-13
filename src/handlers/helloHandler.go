package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello wold"})
}
