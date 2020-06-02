// id UUID not null,
// gauge_id UUID not null,
// gauge_type text not null,
// value double precision not null,
// recorded_date date not null,
// primary key (id),
// constraint metric_gauge_fk foreign key (gauge_id)
// 	references gauge (id)

package model

import (
	"time"

	"github.com/google/uuid"
)

type Metric struct {
	Id           uuid.UUID
	GaugeId      uuid.UUID
	Type         string
	Value        float64
	RecordedDate time.Time
}
