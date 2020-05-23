package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/raptorgandalf/ozark-river-tracker/api/model"
)

type GaugeRepository interface {
	GetAll() (*[]model.Gauge, error)
	Get(uuid.UUID) (*model.Gauge, error)
	Create(river *model.Gauge) error
	Update(river *model.Gauge) error
	Delete(uuid.UUID) error
}

type gaugeRepository struct {
	DB *gorm.DB
}

func GetGaugeRepository(db *gorm.DB) GaugeRepository {
	return &gaugeRepository{
		DB: db,
	}
}

func (r *gaugeRepository) GetAll() (*[]model.Gauge, error) {
	var gauges []model.Gauge

	err := r.DB.Find(&gauges).Error

	return &gauges, err
}

func (r *gaugeRepository) Get(id uuid.UUID) (*model.Gauge, error) {
	var gauge model.Gauge

	err := r.DB.Where("id = ?", id).Take(&gauge).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &gauge, err
}

func (r *gaugeRepository) Create(gauge *model.Gauge) error {
	gauge.Id = uuid.New()

	return r.DB.Create(gauge).Error
}

func (r *gaugeRepository) Update(gauge *model.Gauge) error {
	return r.DB.Save(gauge).Error
}

func (r *gaugeRepository) Delete(id uuid.UUID) error {
	gauge, err := r.Get(id)
	if err != nil {
		return err
	}

	return r.DB.Delete(&gauge).Error
}
