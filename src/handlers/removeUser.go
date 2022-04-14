package handlers

import (
	"example/ginlambda/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveUser(c *gin.Context) {
	id := c.Param("id")

	fmt.Printf("removing user %v", id)

	deletedId, err := repository.DeleteUser(id)

	if err == nil {
		c.JSON(http.StatusAccepted, gin.H{"deletedId": deletedId})
	} else {
		log.Printf("Error removing user")
		c.Error(err)
	}

}
