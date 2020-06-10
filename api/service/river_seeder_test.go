package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/river-folk/ozark-river-tracker/api/service"

	"github.com/river-folk/ozark-river-tracker/api/model"
	"github.com/river-folk/ozark-river-tracker/api/repository"

	"github.com/stretchr/testify/suite"
)

type RiverSeederTestSuite struct {
	suite.Suite
	Db        repository.Database
	RiverSeed model.RiverSeed
}

func (suite *RiverSeederTestSuite) SetupTest() {
	suite.RiverSeed = model.RiverSeed{
		River: model.River{
			Name:      "Current",
			Latitude:  37.2828992,
			Longitude: -91.4103151,
		},
		Gauges: []model.Gauge{
			model.Gauge{
				Name:      "Akers Ferry",
				Code:      "07064533",
				Latitude:  37.3756944,
				Longitude: -91.5528056,
			},
		},
	}

	connection, err := repository.GetConnection()
	if err != nil {
		panic(err)
	}

	connection = connection.Begin()
	suite.Db = repository.GetDatabaseForConnection(connection)
}

func (suite *RiverSeederTestSuite) TearDownTest() {
	suite.Db.Connection.Rollback()
}

func TestRiverSeederTestSuite(t *testing.T) {
	suite.Run(t, new(RiverSeederTestSuite))
}

func (suite *RiverSeederTestSuite) TestSeedRiver() {
	err := service.SeedRiver(suite.RiverSeed, suite.Db)

	assert.Nil(suite.T(), err)

	rivers, _ := suite.Db.RiverRepo.GetAll()
	gauges, _ := suite.Db.GaugeRepo.GetAll()

	assert.Equal(suite.T(), 1, len(*rivers))
	assert.Equal(suite.T(), 1, len(*gauges))
}
