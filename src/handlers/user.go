package handlers

import (
	"errors"
	"example/ginlambda/repository"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	if id := c.Param("id"); id != "" {
		log.Printf("Getting user %v", id)

		userItem, err := repository.GetUser(id)

		if err == nil {
			format := "2006-01-02"
			//Return user
			dob, err := time.Parse(format, userItem.DateOfBirth)
			if err == nil {
				user := UserResponse{
					UserId:      userItem.UserId,
					FirstName:   userItem.FirstName,
					LastName:    userItem.LastName,
					UserName:    userItem.UserName,
					DateOfBirth: dob,
				}
				c.JSON(http.StatusOK, gin.H{"user": user})
			} else {
				log.Printf("Error parsing user dob")
				c.Error(err)
			}
		} else {
			// Error getting user
			log.Printf("Error getting user")
			c.Error(err)
		}
	} else {
		log.Printf("Provided ud field is blank")
		c.Error(errors.New("Id cannot be blank"))
	}

}
