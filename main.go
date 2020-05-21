package main

import (
	"encoding/json"
	"fmt"

	"github.com/raptorgandalf/ozark-river-tracker/pkg/usgs"
)

func main() {
	// timeSeries := usgs.TimeSeriesResponse{
	// 	Name: "Test",
	// }

	// fmt.Println(timeSeries)

	fmt.Println("Ozark river tracker!")

	test, err := usgs.GetSites()
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(test)

	var timeSeries usgs.TimeSeriesResponse
	err = json.Unmarshal(test, &timeSeries)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(timeSeries)
}
