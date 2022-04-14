package handlers

import (
	"example/ginlambda/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func AddUser(c *gin.Context) {

	var newUser AddUserPayload
	err := c.Bind(&newUser)

	if err == nil {
		log.Printf("parsed payload: %+v\n", newUser)
		c.JSON(http.StatusCreated, gin.H{"userId": "029381"})
	} else {
		c.Error(err)
	}
}

func addUser(newUser AddUserPayload) (string, error) {
	uuid := uuid.NewV4().String()
	modelTypeAndId := "User#" + uuid
	item := repository.UserItem{
		UserId:         uuid,
		ModelTypeAndId: modelTypeAndId,
		FirstName:      newUser.FirstName,
		LastName:       newUser.LastName,
		UserName:       newUser.UserName,
		DateOfBirth:    newUser.DateOfBirth.String(),
	}

	createdId, err := repository.AddUser(item)

	if err != nil {
		log.Printf("Error adding user")
		return "", err
	}

	log.Printf("Successfully added user with id %v", uuid)
	return createdId, nil
}
