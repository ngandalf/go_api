package main

import (
	"github.com/ngandalf/go_api/config"
	"github.com/ngandalf/go_api/models"
	"log"
	"net/http"
	"fmt"
)

func main() {
	fmt.Print("Start api")
	config.DatabaseInit()
	router := InitializeRouter()

	// Populate database
	models.NewCar(&models.Car{Manufacturer: "citroen", Design: "ds3", Style: "sport", Doors: 4})

	log.Fatal(http.ListenAndServe(":8081", router))
}
