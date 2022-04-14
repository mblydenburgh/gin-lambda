package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveUser(c *gin.Context) {
	id := c.Param("id")

	fmt.Printf("removing user %v", id)

	c.JSON(http.StatusAccepted, gin.H{"userId": id})
}
