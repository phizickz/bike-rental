package rental

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func(r *Repository) FindRental(id int) (*Rental, error) {
	var rental Rental
	if err := r.db.First(&rental, id).Error; err != nil {
		return nil, err
	}
	return &rental, nil
}

func (r *Repository) CreateRental(rental *Rental) error {
	return r.db.Create(rental).Error
}

func (r *Repository) SaveRental(rental *Rental) error {
	return r.db.Save(rental).Error
}

func (r *Repository) GetAllRentals() ([]Rental, error) {
	var rentals []Rental
	if err := r.db.Find(&rentals).Error; err != nil {
		return nil, err
	}
	return rentals, nil
}