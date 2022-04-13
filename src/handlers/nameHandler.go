package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NameHandler(c *gin.Context) {
	name := c.Param("name")

	c.String(http.StatusOK, "Hello %s", name)
}
