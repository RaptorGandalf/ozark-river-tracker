package service

import (
	"github.com/teamship-studios/ozark-river-tracker/api/model"
	"github.com/teamship-studios/ozark-river-tracker/api/repository"
)

func SeedRiver(riverSeed model.RiverSeed, db repository.Database) error {

	// TODO: Skip river if it already exists
	river := riverSeed.River

	err := db.RiverRepo.Create(&river)
	if err != nil {
		return err
	}

	// TODO: Skip gauge if it already exists
	for _, gauge := range riverSeed.Gauges {
		gauge.RiverId = river.Id
		err := db.GaugeRepo.Create(&gauge)
		if err != nil {
			return err
		}
	}

	return nil
}
