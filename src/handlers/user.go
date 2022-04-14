package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {

	id := c.Param("id")
	fmt.Printf("Getting user %v", id)

	user := UserResponse{
		UserId:      id,
		FirstName:   "Bob",
		LastName:    "Bobbington",
		UserName:    "bbobby00",
		DateOfBirth: time.Date(1989, time.Month(2), 19, 0, 0, 0, 0, time.UTC),
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
