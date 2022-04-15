package handlers

import (
	"example/ginlambda/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RemoveCar(c *gin.Context) {
	vin := c.Param("vin")
	log.Printf("removing car %v", vin)

	deletedVin, err := repository.DeleteCar(vin)
	if err == nil {
		c.JSON(http.StatusAccepted, gin.H{"deletedVin": deletedVin})
	} else {
		log.Printf("Error removing car")
		c.Error(err)
	}
}
