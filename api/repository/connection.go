package repository

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	Connection *gorm.DB
	RiverRepo  RiverRepository
}

func GetDatabase() (Database, error) {
	connection, err := GetConnection()
	if err != nil {
		return Database{}, err
	}

	database := buildDatabase(connection)

	return database, nil
}

func GetDatabaseForConnection(connection *gorm.DB) Database {
	return buildDatabase(connection)
}

func GetConnection() (*gorm.DB, error) {
	connectionString := os.Getenv("PG_CONN")

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(50)

	return db, nil
}

func buildDatabase(connection *gorm.DB) Database {
	return Database{
		Connection: connection,
		RiverRepo:  GetRiverRepository(connection),
	}
}