package main

import (
	"fmt"
	"os"

	"github.com/teamship-studios/ozark-river-tracker/pkg/usgs"
)

func main() {
	gaugeCode := os.Args[1]

	response, err := usgs.ReadGauge(gaugeCode, []string{usgs.GaugeHeight, usgs.Discharge, "91110"})
	if err != nil {
		panic(err)
	}

	discharge, err := response.GetMostRecentDischarge()
	if err != nil {
		panic(err)
	}

	height, err := response.GetMostRecentGaugeHeight()
	if err != nil {
		panic(err)
	}

	latitude, longitude, err := response.GetCoordinates()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Discharge: %f\n", discharge)
	fmt.Printf("Height: %f\n", height)
	fmt.Printf("Latitude: %f\n", latitude)
	fmt.Printf("Longitude: %f\n", longitude)
	fmt.Printf("%f,%f,%f,%f", discharge, height, latitude, longitude)
}
