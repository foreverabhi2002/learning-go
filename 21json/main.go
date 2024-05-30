package main

import (
	"encoding/json"
	"fmt"
)

type bike struct {
	Name             string `json:"bikeName"`
	Engine           string `json:"bikeEngine"`
	Brand            string `json:"bikeBrand"`
	ManufacturedYear int
	FuelType         string
	NumberPlate      string   `json:"-"`
	Tags             []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON")
	EncodeJson()
}

func EncodeJson() {
	bikes := []bike{
		{"Meteor Fireball Matte Green", "350cc", "Royal Enfield", 2023, "Petrol", "JHO5DN5525", []string{"cruiser", "bullet"}},
		{"Activa 5g", "110cc", "Honda", 2018, "Petrol", "JHO5CE4669", []string{"scooter"}},
		{"Shine", "125cc", "Honda", 2016, "Petrol", "UP80DA4280", nil},
	}
	// put this data as json data
	finalJson, err := json.MarshalIndent(bikes, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
