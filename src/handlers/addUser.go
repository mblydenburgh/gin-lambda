package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var newUser AddUserPayload
	c.Bind(&newUser)

	fmt.Printf("payload: %v", newUser)

	c.JSON(http.StatusCreated, gin.H{"userId": "029381"})
}
