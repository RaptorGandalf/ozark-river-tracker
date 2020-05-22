package repository_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/raptorgandalf/ozark-river-tracker/api/model"
	"github.com/raptorgandalf/ozark-river-tracker/api/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RiverTestSuite struct {
	suite.Suite
	Db    repository.Database
	River model.River
}

func (suite *RiverTestSuite) SetupTest() {
	suite.Db, _ = repository.GetDatabase()

	suite.River = model.River{
		Name:      "Current",
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

func (suite *RiverTestSuite) TearDownTest() {
	suite.Db.Connection.Rollback()
}

func TestRiverTestSuite(t *testing.T) {
	suite.Run(t, new(RiverTestSuite))
}

func (suite *RiverTestSuite) TestGetAll() {
	riverA := suite.River
	riverB := suite.River

	suite.Db.RiverRepo.Create(&riverA)
	suite.Db.RiverRepo.Create(&riverB)

	result, err := suite.Db.RiverRepo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(*result))
}

func (suite *RiverTestSuite) TestGet() {
	river := suite.River

	suite.Db.RiverRepo.Create(&river)

	result, err := suite.Db.RiverRepo.Get(river.Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), river.Id, result.Id)
}

func (suite *RiverTestSuite) TestGetNotFound() {
	result, err := suite.Db.RiverRepo.Get(uuid.New())

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *RiverTestSuite) TestCreate() {
	river := suite.River

	err := suite.Db.RiverRepo.Create(&river)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.RiverRepo.Get(river.Id)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), river.Id, actual.Id)
	assert.Equal(suite.T(), river.Name, actual.Name)
	assert.Equal(suite.T(), river.Latitude, actual.Latitude)
	assert.Equal(suite.T(), river.Longitude, actual.Longitude)
}

func (suite *RiverTestSuite) TestUpdate() {
	river := suite.River

	suite.Db.RiverRepo.Create(&river)

	river.Name = "Jacks Fork"

	err := suite.Db.RiverRepo.Update(&river)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.RiverRepo.Get(river.Id)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), "Jacks Fork", actual.Name)
}

func (suite *RiverTestSuite) TestDelete() {
	river := suite.River

	suite.Db.RiverRepo.Create(&river)

	err := suite.Db.RiverRepo.Delete(river.Id)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.RiverRepo.Get(river.Id)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), actual)
}
