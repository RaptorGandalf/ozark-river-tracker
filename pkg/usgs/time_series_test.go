package usgs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/raptorgandalf/ozark-river-tracker/pkg/usgs"
)

type TimeSeriesTestSuite struct {
	suite.Suite
	TimeSeriesResponse usgs.TimeSeriesResponse
}

func (suite *TimeSeriesTestSuite) SetupTest() {
	gageVariableCode := usgs.VariableCode{
		Value: "00065",
	}

	gageValues := []usgs.TimeSeriesValue{
		usgs.TimeSeriesValue{
			Value:    "140",
			DateTime: "2020-05-20T14:30:00.000-05:00",
		},
		usgs.TimeSeriesValue{
			Value:    "200",
			DateTime: "2020-05-20T15:30:00.000-05:00",
		},
	}

	var gageTimeSeries usgs.TimeSeries

	gageTimeSeries.Variable.VariableCode = []usgs.VariableCode{gageVariableCode}
	gageTimeSeries.Values = []usgs.Value{usgs.Value{Value: gageValues}}

	dischargeVariableCode := usgs.VariableCode{
		Value: "00060",
	}

	dischargeValues := []usgs.TimeSeriesValue{
		usgs.TimeSeriesValue{
			Value:    "1000",
			DateTime: "2020-05-20T14:30:00.000-05:00",
		},
		usgs.TimeSeriesValue{
			Value:    "2000",
			DateTime: "2020-05-20T15:30:00.000-05:00",
		},
	}

	var dischargeTimeSeries usgs.TimeSeries

	dischargeTimeSeries.Variable.VariableCode = []usgs.VariableCode{dischargeVariableCode}
	dischargeTimeSeries.Values = []usgs.Value{usgs.Value{Value: dischargeValues}}

	temperatureVariableCode := usgs.VariableCode{
		Value: "00010",
	}

	temperatureValues := []usgs.TimeSeriesValue{
		usgs.TimeSeriesValue{
			Value:    "50",
			DateTime: "2020-05-20T14:30:00.000-05:00",
		},
		usgs.TimeSeriesValue{
			Value:    "60",
			DateTime: "2020-05-20T15:30:00.000-05:00",
		},
	}

	var temperatureTimeSeries usgs.TimeSeries

	temperatureTimeSeries.Variable.VariableCode = []usgs.VariableCode{temperatureVariableCode}
	temperatureTimeSeries.Values = []usgs.Value{usgs.Value{Value: temperatureValues}}

	timeSeries := []usgs.TimeSeries{gageTimeSeries, dischargeTimeSeries, temperatureTimeSeries}

	suite.TimeSeriesResponse.Value.TimeSeries = timeSeries
}

func TestTimeSeriesTestSuite(t *testing.T) {
	suite.Run(t, new(TimeSeriesTestSuite))
}

func (suite *TimeSeriesTestSuite) TestGetMostRecentGageHeight() {
	result, err := suite.TimeSeriesResponse.GetMostRecentGageHeight()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(200), result)
}

func (suite *TimeSeriesTestSuite) TestGetMostRecentDischarge() {
	result, err := suite.TimeSeriesResponse.GetMostRecentDischarge()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(2000), result)
}

func (suite *TimeSeriesTestSuite) TestGetMostRecentWaterTemperature() {
	result, err := suite.TimeSeriesResponse.GetMostRecentWaterTemperature()

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float64(60), result)
}
