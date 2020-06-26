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

	converted, err := strconv.Atoi(os.Getenv("GAUGE_READ_INTERVAL"))
	if err != nil {
		gaugeReadInterval = converted
	}

	Config = Configuration{
		PostgressConnection: os.Getenv("PG_CONN"),
		MigrateOnStartup:    os.Getenv("MIGRATE_ON_STARTUP") == "true",
		GaugeReadInterval:   gaugeReadInterval,
	}
}
func init() {
	metricDeleteDays := 120

	converted, err := strconv.Atoi(os.Getenv("METRIC_DELETE_DAYS"))
	if err != nil {
		metricDeleteDays = converted
	}

	Config = Configuration{
		PostgressConnection: os.Getenv("PG_CONN"),
		MigrateOnStartup:    os.Getenv("MIGRATE_ON_STARTUP") == "true",
		MetricDeleteDays:    metricDeleteDays,
	}
}
