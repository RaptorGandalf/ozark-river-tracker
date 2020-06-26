package configuration

import (
	"os"
	"strconv"
)

type Configuration struct {
	PostgressConnection string
	MigrateOnStartup    bool
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
		MigrateOnStartup:    os.Getenv("MIGRATE_ON_STARTUP") == "true",
		GaugeReadInterval:   gaugeReadInterval,
		MetricDeleteDays:    metricDeleteDays,
	}
}
func init() {
	metricDeleteDays := 120

	converted, err := strconv.Atoi(os.Getenv("METRIC_DELETE_DAYS"))
	if err == nil {
		metricDeleteDays = converted
	}

	Config = Configuration{
		PostgressConnection: os.Getenv("PG_CONN"),
		MigrateOnStartup:    os.Getenv("MIGRATE_ON_STARTUP") == "true",
		MetricDeleteDays:    metricDeleteDays,
	}
}
