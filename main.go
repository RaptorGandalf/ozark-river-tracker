package main

import (
	"fmt"
	"time"

	"github.com/river-folk/ozark-river-tracker/api/service"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/river-folk/ozark-river-tracker/api/repository"
	"github.com/river-folk/ozark-river-tracker/api/router"
)

func main() {
	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Every(15).Minutes().Do(func() {
		db, err := repository.GetDatabase()
		if err != nil {
			fmt.Println(err)
			return
		}

		service.ReadGauges(db)
	})

	scheduler.StartAsync()

	connection, err := repository.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	http := gin.Default()

	router.Setup(http, connection)

	err = http.Run("localhost:80")
	if err != nil {
		fmt.Println(err)
	}
}
