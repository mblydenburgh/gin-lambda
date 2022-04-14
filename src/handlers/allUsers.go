package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AllUsers(c *gin.Context) {
	user := UserResponse{
		UserId:      "1234",
		FirstName:   "Bob",
		LastName:    "Bobbington",
		UserName:    "bbobby00",
		DateOfBirth: time.Date(1989, time.Month(2), 19, 0, 0, 0, 0, time.UTC),
	}
	var users = []UserResponse{
		user,
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
