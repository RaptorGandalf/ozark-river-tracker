package repository_test

import (
	"testing"

	"github.com/teamship-studios/ozark-river-tracker/api/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConnectionTestSuite struct {
	suite.Suite
}

func TestConnectionTestSuite(t *testing.T) {
	suite.Run(t, new(ConnectionTestSuite))
}

func (suite *ConnectionTestSuite) TestGetDatabase() {
	_, err := repository.GetDatabase()

	assert.Nil(suite.T(), err)
}

func (suite *ConnectionTestSuite) TestGetDatabaseForConnection() {
	result, err := repository.GetConnection()

	assert.Nil(suite.T(), err)

	database := repository.GetDatabaseForConnection(result)

	assert.NotNil(suite.T(), database)
}
