package repository_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/raptorgandalf/ozark-river-tracker/api/model"
	"github.com/raptorgandalf/ozark-river-tracker/api/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GaugeTestSuite struct {
	suite.Suite
	Db    repository.Database
	River model.River
	Gauge model.Gauge
}

func (suite *GaugeTestSuite) SetupTest() {
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
}

func (suite *GaugeTestSuite) TearDownTest() {
	suite.Db.Connection.Rollback()
}

func TestGaugeTestSuite(t *testing.T) {
	suite.Run(t, new(GaugeTestSuite))
}

func (suite *GaugeTestSuite) TestGetAll() {
	riverA := suite.River
	riverB := suite.River
	gaugeA := suite.Gauge
	gaugeB := suite.Gauge

	suite.Db.RiverRepo.Create(&riverA)
	suite.Db.RiverRepo.Create(&riverB)

	gaugeA.RiverId = riverA.Id
	gaugeB.RiverId = riverB.Id

	suite.Db.GaugeRepo.Create(&gaugeA)
	suite.Db.GaugeRepo.Create(&gaugeB)

	result, err := suite.Db.GaugeRepo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(*result))
}

func (suite *GaugeTestSuite) TestGet() {
	river := suite.River
	gauge := suite.Gauge

	suite.Db.RiverRepo.Create(&river)

	gauge.RiverId = river.Id

	suite.Db.GaugeRepo.Create(&gauge)

	result, err := suite.Db.GaugeRepo.Get(gauge.Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), gauge.Id, result.Id)
}

func (suite *GaugeTestSuite) TestGetNotFound() {
	result, err := suite.Db.GaugeRepo.Get(uuid.New())

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *GaugeTestSuite) TestCreate() {
	river := suite.River
	gauge := suite.Gauge

	suite.Db.RiverRepo.Create(&river)

	gauge.RiverId = river.Id

	err := suite.Db.GaugeRepo.Create(&gauge)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.GaugeRepo.Get(gauge.Id)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), gauge.Id, actual.Id)
	assert.Equal(suite.T(), gauge.Name, actual.Name)
	assert.Equal(suite.T(), gauge.Code, actual.Code)
	assert.Equal(suite.T(), gauge.RiverId, actual.RiverId)
	assert.Equal(suite.T(), gauge.Latitude, actual.Latitude)
	assert.Equal(suite.T(), gauge.Longitude, actual.Longitude)
}

func (suite *GaugeTestSuite) TestUpdate() {
	river := suite.River
	gauge := suite.Gauge

	suite.Db.RiverRepo.Create(&river)

	gauge.RiverId = river.Id

	suite.Db.GaugeRepo.Create(&gauge)

	gauge.Name = "This here be a gauge1"

	err := suite.Db.GaugeRepo.Update(&gauge)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.GaugeRepo.Get(gauge.Id)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), "This here be a gauge1", actual.Name)
}

func (suite *GaugeTestSuite) TestDelete() {
	river := suite.River
	gauge := suite.Gauge

	suite.Db.RiverRepo.Create(&river)

	gauge.RiverId = river.Id

	suite.Db.GaugeRepo.Create(&gauge)

	err := suite.Db.GaugeRepo.Delete(gauge.Id)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.GaugeRepo.Get(gauge.Id)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), actual)
}
