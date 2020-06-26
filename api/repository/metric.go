package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/river-folk/ozark-river-tracker/api/model"
)

type MetricRepository interface {
	GetAll() (*[]model.Metric, error)
	GetGaugeMetrics(gaugeId uuid.UUID) (*[]model.Metric, error)
	Get(uuid.UUID) (*model.Metric, error)
	Create(Metric *model.Metric) error
	Update(Metric *model.Metric) error
	Delete(uuid.UUID) error
	DeleteOldMetrics() error
}

type metricRepository struct {
	DB *gorm.DB
}

func GetMetricRepository(db *gorm.DB) MetricRepository {
	return &metricRepository{
		DB: db,
	}
}

func (r *metricRepository) GetAll() (*[]model.Metric, error) {
	var metrics []model.Metric

	err := r.DB.Find(&metrics).Error

	return &metrics, err
}

func (r *metricRepository) GetGaugeMetrics(gaugeId uuid.UUID) (*[]model.Metric, error) {
	var metrics []model.Metric

	err := r.DB.Where("gauge_id = ?", gaugeId).Find(&metrics).Error

	return &metrics, err
}

func (r *metricRepository) Get(id uuid.UUID) (*model.Metric, error) {
	var metric model.Metric

	err := r.DB.Where("id = ?", id).Take(&metric).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &metric, err
}

func (r *metricRepository) Create(metric *model.Metric) error {
	metric.Id = uuid.New()

	return r.DB.Create(metric).Error
}

func (r *metricRepository) Update(metric *model.Metric) error {
	return r.DB.Save(metric).Error
}

func (r *metricRepository) Delete(id uuid.UUID) error {
	metric, err := r.Get(id)
	if err != nil {
		return err
	}

	return r.DB.Delete(&metric).Error
}

func (r *metricRepository) DeleteOldMetrics() error {
	var cutoff = time.Now().AddDate(0, 0, -120)
	return r.DB.Where("RecordedDate <= ?", cutoff).Delete(model.Metric{}).Error

}
