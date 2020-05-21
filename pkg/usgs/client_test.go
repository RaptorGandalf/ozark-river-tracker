package usgs_test

import (
	"testing"

	"github.com/raptorgandalf/ozark-river-tracker/pkg/usgs"
	gock "gopkg.in/h2non/gock.v1"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (suite *ClientTestSuite) SetupTest() {
	mock := usgs.NewMock()
	mock.Start()
}

func (suite *ClientTestSuite) TearDownTest() {
	gock.Off()
}

func (suite *ClientTestSuite) TestGetData() {
	result, err := usgs.GetData([]string{"07064533"}, []string{usgs.GageHeight, usgs.Discharge})

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "ns720:timeSeriesResponseType", result.Name)
}
