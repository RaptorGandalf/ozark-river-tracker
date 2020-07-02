package service

import (
	"github.com/teamship-studios/ozark-river-tracker/api/model"
	"github.com/teamship-studios/ozark-river-tracker/api/repository"
)

func SeedRiver(riverSeed model.RiverSeed, db repository.Database) error {
	river := riverSeed.River

	err := db.RiverRepo.Create(&river)
	if err != nil {
		return err
	}

	for _, gauge := range riverSeed.Gauges {
		gauge.RiverId = river.Id
		err := db.GaugeRepo.Create(&gauge)
		if err != nil {
			return err
		}
	}

	return nil
}
