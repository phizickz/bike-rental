package customer

import (
	"fmt"
	"sync"
)

type Repository struct {
	customers map[int]*Customer
	mu        sync.Mutex
	nextID    int
}

func NewRepository() *Repository {
	return &Repository{
		customers: make(map[int]*Customer),
		nextID:    1,
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
	r.mu.Lock()
	defer r.mu.Unlock()
	customer.ID = r.nextID
	r.customers[customer.ID] = customer
	r.nextID++
}

func (r *Repository) PrintCustomers() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, customer := range r.customers {
		fmt.Println(customer)
	}
}
