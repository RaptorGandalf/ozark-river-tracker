package main

import (
	"fmt"

	"github.com/raptorgandalf/ozark-river-tracker/api/repository"
	"github.com/raptorgandalf/ozark-river-tracker/pkg/usgs"
)

func main() {
	fmt.Println("Ozark river tracker!")

	test, err := usgs.GetData([]string{"07064533"}, []string{usgs.GageHeight, usgs.Discharge})
	if err != nil {
		fmt.Println(err)
		return
	}

	discharge, _ := test.GetMostRecentDischarge()
	height, _ := test.GetMostRecentGageHeight()

	fmt.Println(discharge)
	fmt.Println(height)

	_, err = repository.GetDatabase()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Got database!")
	}
}
