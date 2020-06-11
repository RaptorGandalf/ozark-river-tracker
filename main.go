package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/river-folk/ozark-river-tracker/configuration"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	"github.com/river-folk/ozark-river-tracker/api/service"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/river-folk/ozark-river-tracker/api/repository"
	"github.com/river-folk/ozark-river-tracker/api/router"
)

func main() {
	var connection *gorm.DB

	for {
		con, err := repository.GetConnection()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Retrying in 10 seconds.")
			time.Sleep(time.Second * 10)
		} else {
			connection = con
			break
		}
	}

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

	files, err := ioutil.ReadDir("db/migrations")
	if err != nil {
		panic(err)
	}
	fmt.Println(files)

	migration, err := migrate.New("file://db/migrations/", configuration.Config.PostgressConnection)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
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
