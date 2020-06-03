package main

import (
	"fmt"

	"github.com/river-folk/ozark-river-tracker/api/service"
)

func main() {
	fmt.Println("Reading gauges...")

	service.ReadGauges()

	fmt.Println("Done.")

	fmt.Println("Gauge read complete.")
}
