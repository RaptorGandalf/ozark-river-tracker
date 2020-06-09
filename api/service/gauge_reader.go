// TODO: write tests

package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/river-folk/ozark-river-tracker/api/model"
	"github.com/river-folk/ozark-river-tracker/api/repository"
	"github.com/river-folk/ozark-river-tracker/pkg/usgs"
)

func ReadGauges(db repository.Database) {
	fmt.Println("Loading rivers...")

	rivers, err := db.RiverRepo.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Rivers loaded.")

	for _, river := range *rivers {
		fmt.Printf("Loading gauges for %s river.\n", river.Name)

		gauges, err := db.GaugeRepo.GetRiverGauges(river.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Gauges loaded.")

		for _, gauge := range *gauges {

			fmt.Printf("Reading %s gauge on %s river.\n", gauge.Name, river.Name)

			height, discharge, err := readGauge(gauge)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("Creating metrics for %s gauge on %s river.\n", gauge.Name, river.Name)

			createMetrics(gauge.Id, height, discharge, db)
		}
	}
}

func readGauge(gauge model.Gauge) (float64, float64, error) {
	gaugeData, err := usgs.ReadGauge(gauge.Code, []string{usgs.GaugeHeight, usgs.Discharge})
	if err != nil {
		return -1, -1, err
	}

	height, err := gaugeData.GetMostRecentGaugeHeight()
	if err != nil {
		height = -1
	}

	discharge, err := gaugeData.GetMostRecentDischarge()
	if err != nil {
		discharge = -1
	}

	return height, discharge, nil
}

func createMetrics(gaugeId uuid.UUID, height, discharge float64, db repository.Database) {
	createMetric(gaugeId, height, "height", db)
	createMetric(gaugeId, discharge, "discharge", db)
}

func createMetric(gaugeId uuid.UUID, value float64, metricType string, db repository.Database) {
	fmt.Println(value)

	if value < 0 {
		return
	}

	metric := model.Metric{
		GaugeId:      gaugeId,
		Type:         metricType,
		Value:        value,
		RecordedDate: time.Now(),
	}

	db.MetricRepo.Create(&metric)
}
