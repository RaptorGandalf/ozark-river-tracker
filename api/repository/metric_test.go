package repository_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/river-folk/ozark-river-tracker/api/model"
	"github.com/river-folk/ozark-river-tracker/api/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MetricTestSuite struct {
	suite.Suite
	Db     repository.Database
	River  model.River
	Gauge  model.Gauge
	Metric model.Metric
}

func (suite *MetricTestSuite) SetupTest() {
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

	suite.Metric = model.Metric{
		Type:         "Height",
		Value:        2.40,
		RecordedDate: time.Now(),
	}

	connection, err := repository.GetConnection()
	if err != nil {
		panic(err)
	}

	connection = connection.Begin()
	suite.Db = repository.GetDatabaseForConnection(connection)
}

func (suite *MetricTestSuite) TearDownTest() {
	suite.Db.Connection.Rollback()
}

func TestMetricTestSuite(t *testing.T) {
	suite.Run(t, new(MetricTestSuite))
}

func (suite *MetricTestSuite) TestGetAll() {
	river := suite.River
	suite.Db.RiverRepo.Create(&river)

	gauge := suite.Gauge
	gauge.RiverId = river.Id
	suite.Db.GaugeRepo.Create(&gauge)

	metric := suite.Metric
	metric.GaugeId = gauge.Id
	suite.Db.MetricRepo.Create(&metric)

	result, err := suite.Db.MetricRepo.GetAll()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, len(*result))
}

func (suite *MetricTestSuite) TestGetGaugeMetrics() {
	river := suite.River
	suite.Db.RiverRepo.Create(&river)

	gauge := suite.Gauge
	gauge.RiverId = river.Id
	suite.Db.GaugeRepo.Create(&gauge)

	metricA := suite.Metric
	metricA.GaugeId = gauge.Id
	suite.Db.MetricRepo.Create(&metricA)

	metricB := suite.Metric
	metricB.GaugeId = gauge.Id
	suite.Db.MetricRepo.Create(&metricB)

	result, err := suite.Db.MetricRepo.GetGaugeMetrics(gauge.Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, len(*result))
}

func (suite *MetricTestSuite) TestGet() {
	river := suite.River
	suite.Db.RiverRepo.Create(&river)

	gauge := suite.Gauge
	gauge.RiverId = river.Id
	suite.Db.GaugeRepo.Create(&gauge)

	metric := suite.Metric
	metric.GaugeId = gauge.Id
	suite.Db.MetricRepo.Create(&metric)

	result, err := suite.Db.MetricRepo.Get(metric.Id)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), metric.Id, result.Id)
}

func (suite *MetricTestSuite) TestGetNotFound() {
	result, err := suite.Db.MetricRepo.Get(uuid.New())

	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *MetricTestSuite) TestCreate() {
	river := suite.River
	suite.Db.RiverRepo.Create(&river)

	gauge := suite.Gauge
	gauge.RiverId = river.Id
	suite.Db.GaugeRepo.Create(&gauge)

	metric := suite.Metric
	metric.GaugeId = gauge.Id
	fmt.Println("=============================")
	fmt.Println(metric)
	fmt.Println("=============================")
	err := suite.Db.MetricRepo.Create(&metric)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.MetricRepo.Get(metric.Id)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), metric.Id, actual.Id)
	assert.Equal(suite.T(), metric.GaugeId, actual.GaugeId)
	assert.Equal(suite.T(), metric.Value, actual.Value)
}

func (suite *MetricTestSuite) TestUpdate() {
	river := suite.River
	suite.Db.RiverRepo.Create(&river)

	gauge := suite.Gauge
	gauge.RiverId = river.Id
	suite.Db.GaugeRepo.Create(&gauge)

	metric := suite.Metric
	metric.GaugeId = gauge.Id
	suite.Db.MetricRepo.Create(&metric)

	metric.Value = 3.4

	err := suite.Db.MetricRepo.Update(&metric)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.MetricRepo.Get(metric.Id)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), 3.4, actual.Value)
}

func (suite *MetricTestSuite) TestDelete() {
	river := suite.River
	suite.Db.RiverRepo.Create(&river)

	gauge := suite.Gauge
	gauge.RiverId = river.Id
	suite.Db.GaugeRepo.Create(&gauge)

	metric := suite.Metric
	metric.GaugeId = gauge.Id
	suite.Db.MetricRepo.Create(&metric)

	err := suite.Db.MetricRepo.Delete(metric.Id)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.MetricRepo.Get(metric.Id)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), actual)
}

func (suite.MetricTestSuite) TestDeleteOldMetrics() {
	river := suite.River
	suite.Db.RiverRepo.Create(&river)

	gauge := suite.Gauge
	gauge.RiverId = river.Id
	suite.Db.GaugeRepo.Create(&gauge)

	metric := suite.Metric
	metric.GaugeId = gauge.Id
	suite.Db.MetricRepo.Create(&metric)

	err := suite.Db.MetricRepo.DeleteOldMetrics(metric.Id)

	metric.RecordedDate = time.Now().AddDate(0, 0, -120)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.MetricRepo.Get(metric.Id)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), actual)

	metric.RecordedDate = time.Now().AddDate(0, 0, -80)

	assert.Nil(suite.T(), err)

	actual, err := suite.Db.MetricRepo.Get(metric.Id)
	assert.Nil(suite.T(), err)
	assert.Nil(suite.T(), actual)
}
