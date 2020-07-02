package main

import (
	"fmt"

	"github.com/teamship-studios/ozark-river-tracker/api/repository"
	"github.com/teamship-studios/ozark-river-tracker/api/service"
)

func main() {
	fmt.Println("Reading gauges...")

	db, err := repository.GetDatabase()
	if err != nil {
		panic(err)
	}

	service.ReadGauges(db)

	fmt.Println("Done.")

	fmt.Println("Gauge read complete.")
}
