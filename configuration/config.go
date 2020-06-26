package configuration

import (
	"os"
	"strconv"

	"github.com/rollbar/rollbar-go"
)

type Configuration struct {
	PostgressConnection string
	GaugeReadInterval   int
	MetricDeleteDays    int
}

var Config Configuration

func init() {
	gaugeReadInterval := 15

	gaugeConverted, err := strconv.Atoi(os.Getenv("GAUGE_READ_INTERVAL"))
	if err == nil {
		gaugeReadInterval = gaugeConverted
	}

	metricDeleteDays := 120

	deleteConverted, err := strconv.Atoi(os.Getenv("METRIC_DELETE_DAYS"))
	if err == nil {
		metricDeleteDays = deleteConverted
	}

	Config = Configuration{
		PostgressConnection: os.Getenv("PG_CONN"),
		GaugeReadInterval:   gaugeReadInterval,
		MetricDeleteDays:    metricDeleteDays,
	}

	ortEnv := os.Getenv("ORT_ENV")
	if ortEnv == "" {
		ortEnv = "development"
	}

	rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
	rollbar.SetEnvironment(os.Getenv("ORT_ENV"))
	rollbar.SetServerRoot("https://github.com/river-folk/ozark-river-tracker")
}
