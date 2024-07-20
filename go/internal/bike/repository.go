package bike

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB

}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}

}

func (r *Repository) FindBike(id int) (*Bike, error) {
	var bike Bike
	if err := r.db.First(&bike, id).Error; err != nil {
		return nil, err
	}
	return &bike, nil

}

func (r *Repository) SaveBike(bike *Bike) error {
	return r.db.Save(bike).Error

}

func (r *Repository) CreateBike(bike *Bike) error {
	return r.db.Create(bike).Error
}

func (r *Repository) GetAllBikes() ([]Bike, error) {
	var bikes []Bike
	if err := r.db.Find(&bikes).Error; err != nil {
		return nil, err
	}
	return bikes, nil
}
