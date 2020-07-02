package jobs

import (
	"github.com/teamship-studios/ozark-river-tracker/api/repository"
	"github.com/teamship-studios/ozark-river-tracker/api/service"
	"github.com/teamship-studios/ozark-river-tracker/pkg/common"
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

	defer db.Connection.Close()

	err = db.MetricRepo.DeleteOldMetrics()
	if err != nil {
		panic(err)
	}
}
