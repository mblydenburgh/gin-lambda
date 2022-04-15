package handlers

import (
	"example/ginlambda/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCar(c *gin.Context) {

	var newCar AddCarPayload
	err := c.Bind(&newCar)

	// If there is no error from binding incoming json to struct, continue
	if err == nil {
		log.Printf("parsed payload: %+v\n", newCar)

		// Call private method that interacts with ddb
		vin, err := addCar(newCar)

		// If no error, continue
		if err == nil {
			c.JSON(http.StatusCreated, gin.H{"createdVIN": vin})
		} else {
			// Error with addUser method
			c.Error(err)
		}
	} else {
		// Error parsing json
		c.Error(err)
	}

}

func addCar(newCar AddCarPayload) (string, error) {
	modelTypeAndId := "Car#" + newCar.VIN
	item := repository.CarItem{
		UserId:         newCar.UserId,
		ModelTypeAndId: modelTypeAndId,
		Manufacturer:   newCar.Manufacturer,
		Model:          newCar.Model,
		Year:           newCar.Year,
		VehicleType:    newCar.VehicleTyle,
		VIN:            newCar.VIN,
		Color:          newCar.Color,
	}

	createdVIN, err := repository.AddCar(item)

	if err != nil {
		log.Printf("Error adding user")
		return "", err
	}

	log.Printf("Successfully added car with vin %v", createdVIN)
	return createdVIN, nil

}
