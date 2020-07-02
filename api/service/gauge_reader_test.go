package service_test

import (
	"testing"

	"github.com/teamship-studios/ozark-river-tracker/api/service"
	"github.com/teamship-studios/ozark-river-tracker/pkg/usgs"
	"gopkg.in/h2non/gock.v1"

	"github.com/teamship-studios/ozark-river-tracker/api/model"
	"github.com/teamship-studios/ozark-river-tracker/api/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GaugeReaderTestSuite struct {
	suite.Suite
	Db    repository.Database
	River model.River
	Gauge model.Gauge
}

func (suite *GaugeReaderTestSuite) SetupTest() {
	suite.Db, _ = repository.GetDatabase()

	suite.River = model.River{
		Name:      "Current",
		Latitude:  42.0,
		Longitude: 52.0,
	}

	suite.Gauge = model.Gauge{
		Name:      "Gauge1",
		Code:      "12345",
		Latitude:  42.0,
		Longitude: 52.0,
	}

	connection, err := repository.GetConnection()
	if err != nil {
		panic(err)
	}

	connection = connection.Begin()
	suite.Db = repository.GetDatabaseForConnection(connection)

	mock := usgs.NewMock()
	mock.Start()
}

func (suite *GaugeReaderTestSuite) TearDownTest() {
	suite.Db.Connection.Rollback()
	gock.Off()
}

func TestGaugeReaderTestSuite(t *testing.T) {
	suite.Run(t, new(GaugeReaderTestSuite))
}

func (suite *GaugeReaderTestSuite) TestReadGauges() {
	river := suite.River
	err := suite.Db.RiverRepo.Create(&river)
	assert.Nil(suite.T(), err)

	gauge := suite.Gauge
	gauge.RiverId = river.Id
	err = suite.Db.GaugeRepo.Create(&gauge)
	assert.Nil(suite.T(), err)

	service.ReadGauges(suite.Db)

	result, err := suite.Db.MetricRepo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(*result))

	list := *result
	assert.Equal(suite.T(), gauge.Id, list[0].GaugeId)
}
