package customer

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB

}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}

}

func (r *Repository) FindCustomer(id int) (*Customer, error) {

	var customer Customer
	if err := r.db.First(&customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *Repository) CreateCustomer(customer *Customer) error {
	return r.db.Create(customer).Error
}

func (r *Repository) SaveCustomer(customer *Customer) error {
	return r.db.Save(customer).Error
}

func (r *Repository) GetAllCustomers() ([]Customer, error) {
	var customers []Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
