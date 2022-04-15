package main

import (
	"example/ginlambda/handlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway/v2"
	"github.com/gin-gonic/gin"
)

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/users", handlers.AllUsers)
	r.GET("/users/:id", handlers.User)
	r.POST("users", handlers.AddUser)
	r.DELETE("users/:id", handlers.RemoveUser)

	//r.GET("/cars/:vin", handlers.Car)
	//r.POST("cars", handlers.AddCar)
	//r.DELETE("cars/:vin", handlers.RemoveCar)

	return r
}

func main() {
	if inLambda() {
		fmt.Println("running aws lambda in aws")
		log.Fatal(gateway.ListenAndServe(":8080", setupRouter()))
	} else {
		fmt.Println("running aws lambda in local")
		log.Fatal(http.ListenAndServe(":8080", setupRouter()))
	}
}
