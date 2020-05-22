package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raptorgandalf/ozark-river-tracker/api/repository"
	"github.com/raptorgandalf/ozark-river-tracker/api/router"
)

func main() {
	// fmt.Println("Ozark river tracker!")

	// test, err := usgs.GetData([]string{"07064533"}, []string{usgs.GageHeight, usgs.Discharge})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// discharge, _ := test.GetMostRecentDischarge()
	// height, _ := test.GetMostRecentGageHeight()

	// fmt.Println(discharge)
	// fmt.Println(height)

	connection, err := repository.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	http := gin.Default()

	router.Setup(http, connection)

	// TODO Set with ENV var or maybe just use "localhost"
	err = http.Run("127.0.0.1:80")
	if err != nil {
		fmt.Println(err)
	}
}
