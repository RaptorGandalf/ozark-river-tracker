package router_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/teamship-studios/ozark-river-tracker/api/model"
	"github.com/teamship-studios/ozark-river-tracker/api/repository"
	"github.com/teamship-studios/ozark-river-tracker/api/router"
	"github.com/teamship-studios/ozark-river-tracker/pkg/test_utility"

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

	http := gin.Default()
	router.Setup(http, connection)

	test_utility.InitializeTestRouter(http)
}

func (suite *GaugeTestSuite) TearDownTest() {
	suite.Db.Connection.Rollback()
}

func TestGaugeTestSuite(t *testing.T) {
	suite.Run(t, new(GaugeTestSuite))
}

func (suite *GaugeTestSuite) TestGetAll() {
	response := test_utility.PerformRequest("GET", "/api/gauges", "")

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Contains(suite.T(), response.Body.String(), "gauges")
}

func (suite *GaugeTestSuite) TestGetRiverGauges() {
	response := test_utility.PerformRequest("GET", "/api/rivers/f6699607-1aae-4e6a-bc04-92f594c061e5/gauges", "")

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Contains(suite.T(), response.Body.String(), "gauges")
}

func (suite *GaugeTestSuite) TestGet() {
	river := suite.River

	suite.Db.RiverRepo.Create(&river)

	gauge := suite.Gauge

	gauge.RiverId = river.Id

	suite.Db.GaugeRepo.Create(&gauge)

	resp := test_utility.PerformRequest("GET", "/api/gauges/"+gauge.Id.String(), "")

	assert.Equal(suite.T(), http.StatusOK, resp.Code)
	assert.Contains(suite.T(), resp.Body.String(), gauge.Id.String())
}
