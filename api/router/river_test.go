package router_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/raptorgandalf/ozark-river-tracker/api/model"
	"github.com/raptorgandalf/ozark-river-tracker/api/repository"
	"github.com/raptorgandalf/ozark-river-tracker/api/router"
	"github.com/raptorgandalf/ozark-river-tracker/pkg/test_utility"

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

	http := gin.Default()
	router.Setup(http, connection)

	test_utility.InitializeTestRouter(http)
}

func (suite *RiverTestSuite) TearDownTest() {
	suite.Db.Connection.Rollback()
}

func TestRiverTestSuite(t *testing.T) {
	suite.Run(t, new(RiverTestSuite))
}

func (suite *RiverTestSuite) TestGetAll() {
	response := test_utility.PerformRequest("GET", "/api/rivers", "")

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Contains(suite.T(), response.Body.String(), "rivers")
}

func (suite *RiverTestSuite) TestGet() {
	river := suite.River

	suite.Db.RiverRepo.Create(&river)

	resp := test_utility.PerformRequest("GET", "/api/rivers/"+river.Id.String(), "")

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.Contains(suite.T(), resp.Body.String(), river.Id.String())
}
