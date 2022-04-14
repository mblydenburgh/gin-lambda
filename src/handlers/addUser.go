package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var newUser AddUserPayload
	err := c.Bind(&newUser)

	if err == nil {
		fmt.Printf("parsed payload: %+v\n", newUser)
		c.JSON(http.StatusCreated, gin.H{"userId": "029381"})
	}

	c.Error(err)
}
