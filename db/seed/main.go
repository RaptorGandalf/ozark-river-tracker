package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/river-folk/ozark-river-tracker/api/model"
	"github.com/river-folk/ozark-river-tracker/api/repository"
	"github.com/river-folk/ozark-river-tracker/api/service"
)

func main() {
	db, err := repository.GetDatabase()
	if err != nil {
		panic(err)
	}

	riverDir := os.Args[1]

	files, err := ioutil.ReadDir(riverDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(riverDir + file.Name())
		if err != nil {
			fmt.Println(err)
			return
		}

		var riverSeed model.RiverSeed

		err = json.Unmarshal(data, &riverSeed)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = service.SeedRiver(riverSeed, db)
	}
}
