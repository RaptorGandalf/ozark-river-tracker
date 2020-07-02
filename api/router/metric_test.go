package router_test

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/teamship-studios/ozark-river-tracker/api/repository"
	"github.com/teamship-studios/ozark-river-tracker/api/router"
	"github.com/teamship-studios/ozark-river-tracker/pkg/test_utility"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MetricTestSuite struct {
	suite.Suite
	Db repository.Database
}

func (suite *MetricTestSuite) SetupTest() {
	suite.Db, _ = repository.GetDatabase()

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

func (suite *MetricTestSuite) TearDownTest() {
	suite.Db.Connection.Rollback()
}

func TestMetricTestSuite(t *testing.T) {
	suite.Run(t, new(MetricTestSuite))
}

func (suite *MetricTestSuite) TestGetGaugeMetrics() {
	response := test_utility.PerformRequest("GET", "/api/gauges/f6699607-1aae-4e6a-bc04-92f594c061e5/metrics", "")

	assert.Equal(suite.T(), http.StatusOK, response.Code)
	assert.Contains(suite.T(), response.Body.String(), "metrics")
}
