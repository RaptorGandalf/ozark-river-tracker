package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/raptorgandalf/ozark-river-tracker/api/model"
)

type RiverRepository interface {
	GetAll() (*[]model.River, error)
	Get(uuid.UUID) (*model.River, error)
	Create(river *model.River) error
	Update(river *model.River) error
	Delete(uuid.UUID) error
}

type riverRepository struct {
	DB *gorm.DB
}

func GetRiverRepository(db *gorm.DB) RiverRepository {
	return &riverRepository{
		DB: db,
	}
}

func (r *riverRepository) GetAll() (*[]model.River, error) {
	var rivers []model.River

	err := r.DB.Find(&rivers).Error

	return &rivers, err
}

func (r *riverRepository) Get(id uuid.UUID) (*model.River, error) {
	var river model.River

	err := r.DB.Where("id = ?", id).Take(&river).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &river, err
}

func (r *riverRepository) Create(river *model.River) error {
	river.Id = uuid.New()

	return r.DB.Create(river).Error
}

func (r *riverRepository) Update(river *model.River) error {
	return r.DB.Save(river).Error
}

func (r *riverRepository) Delete(id uuid.UUID) error {
	river, err := r.Get(id)
	if err != nil {
		return err
	}

	return r.DB.Delete(&river).Error
}
