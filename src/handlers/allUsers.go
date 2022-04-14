package handlers

import (
	"example/ginlambda/repository"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AllUsers(c *gin.Context) {

	userItems, err := repository.GetAllUsers()

	if err == nil && userItems != nil {
		log.Printf("Returning all users")

		users := make([]UserResponse, len(*userItems))
		for i, userItem := range *userItems {
			dob, _ := time.Parse("2006-01-02", userItem.DateOfBirth)
			user := UserResponse{
				UserId:      userItem.UserId,
				FirstName:   userItem.FirstName,
				LastName:    userItem.LastName,
				UserName:    userItem.UserName,
				DateOfBirth: dob,
			}
			users[i] = user
		}

		c.JSON(http.StatusOK, gin.H{"users": users})
	} else {
		c.Error(err)
	}
}
