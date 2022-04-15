package handlers

import (
	"errors"
	"example/ginlambda/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Car(c *gin.Context) {
	if vin := c.Param("vin"); vin != "" {
		log.Printf("Getting car by vin %v", vin)

		carItem, err := repository.GetCar(vin)

		if err == nil {
			//Return car
			car := CarResponse{
				UserId:       carItem.UserId,
				VIN:          carItem.VIN,
				Manufacturer: carItem.Manufacturer,
				Model:        carItem.Model,
				Year:         carItem.Year,
				VehicleTyle:  carItem.VehicleType,
				Color:        carItem.Color,
			}
			c.JSON(http.StatusOK, gin.H{"car": car})
		} else {
			// Error getting user
			log.Printf("Error getting car")
			c.Error(err)
		}
	} else {
		log.Printf("Provided vin field is blank")
		c.Error(errors.New("VIN cannot be blank"))
	}

}
