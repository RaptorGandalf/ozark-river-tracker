package jobs

import (
	"github.com/river-folk/ozark-river-tracker/api/repository"
	"github.com/river-folk/ozark-river-tracker/api/service"
	"github.com/river-folk/ozark-river-tracker/pkg/common"
)

func PerformReadGauges() {
	defer common.Rescue()

	db, err := repository.GetDatabase()
	if err != nil {
		panic(err)
	}

	service.ReadGauges(db)
}

func PerformCleanMetrics() {
	defer common.Rescue()

	db, err := repository.GetDatabase()
	if err != nil {
		panic(err)
	}

	err = db.MetricRepo.DeleteOldMetrics()
	if err != nil {
		panic(err)
	}
}
