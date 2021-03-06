package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/teamship-studios/ozark-river-tracker/configuration"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/teamship-studios/ozark-river-tracker/api/jobs"
	"github.com/teamship-studios/ozark-river-tracker/api/repository"
	"github.com/teamship-studios/ozark-river-tracker/api/router"
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
		jobs.PerformReadGauges()
	})

	scheduler.Every(1).Days().Do(func() {
		jobs.PerformCleanMetrics()
	})

	scheduler.StartAsync()

	files, err := ioutil.ReadDir("db/migrations")
	if err != nil {
		fmt.Println(err)
		return
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
	http.Use(cors.Default())

	router.Setup(http, connection)

	err = http.Run("0.0.0.0:80")
	if err != nil {
		fmt.Println(err)
	}
}
