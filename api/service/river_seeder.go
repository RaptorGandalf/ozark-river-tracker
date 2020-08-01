package service

import (
	"github.com/teamship-studios/ozark-river-tracker/api/model"
	"github.com/teamship-studios/ozark-river-tracker/api/repository"
)

func SeedRiver(riverSeed model.RiverSeed, db repository.Database) error {

	river := riverSeed.River

	exists, err := db.RiverRepo.GetByName(river.Name)
	if err != nil {
		return err
	}

	if exists != nil {
		return nil
	}

	err = db.RiverRepo.Create(&river)
	if err != nil {
		return err
	}

	gauges, err := db.GaugeRepo.GetRiverGauges(river.Id)
	if err != nil {
		return err
	}

	for _, gauge := range riverSeed.Gauges {

		if gauges != nil {
			if gaugeExists(gauge.Name, *gauges) {
				continue
			}
		}

		gauge.RiverId = river.Id
		err := db.GaugeRepo.Create(&gauge)
		if err != nil {
			return err
		}
	}

	return nil
}

func gaugeExists(name string, gauges []model.Gauge) bool {
	for _, gauge := range gauges {
		if gauge.Name == name {
			return true
		}
	}
	return false
}
