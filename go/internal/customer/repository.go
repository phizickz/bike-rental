package customer

import "fmt"

type Repository struct {
	customers map[int]*Customer
}

func NewRepository() *Repository {
	return &Repository{
		customers: make(map[int]*Customer),
	}
}

func (r *Repository) FindCustomer(id int) (*Customer, error) {
	customer, exists := r.customers[id]
	if !exists {
		return nil, fmt.Errorf("Customer does not exist.")
	}
	return customer, nil
}

func (r *Repository) SaveCustomer(customer *Customer) {
	r.customers[customer.ID] = customer
}

func (r *Repository) GetHighestID() int {
	return len(r.customers)
}