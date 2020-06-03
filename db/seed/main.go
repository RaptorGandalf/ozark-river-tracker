package main

import (
	"github.com/river-folk/ozark-river-tracker/api/model"
	"github.com/river-folk/ozark-river-tracker/api/repository"
)

func main() {
	db, err := repository.GetDatabase()
	if err != nil {
		panic(err)
	}

	// Current River
	currentRiver := model.River{
		Name:      "Current",
		Latitude:  37.2828992,
		Longitude: -91.4103151,
	}

	err = db.RiverRepo.Create(&currentRiver)
	if err != nil {
		panic(err)
	}

	akers := model.Gauge{
		Name:      "Akers Ferry",
		RiverId:   currentRiver.Id,
		Code:      "07064533",
		Latitude:  37.3756944,
		Longitude: -91.5528056,
	}

	err = db.GaugeRepo.Create(&akers)
	if err != nil {
		panic(err)
	}

	powderMill := model.Gauge{
		Name:      "Powder Mill",
		RiverId:   currentRiver.Id,
		Code:      "07066510",
		Latitude:  37.18561389,
		Longitude: -91.1776639,
	}

	err = db.GaugeRepo.Create(&powderMill)
	if err != nil {
		panic(err)
	}

	vanBuren := model.Gauge{
		Name:      "Van Buren",
		RiverId:   currentRiver.Id,
		Code:      "07067000",
		Latitude:  36.99138889,
		Longitude: -91.0135,
	}

	err = db.GaugeRepo.Create(&vanBuren)
	if err != nil {
		panic(err)
	}
}
