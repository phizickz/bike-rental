package bike

import (
	"fmt"
	"sync"
)

type Repository struct {
	bikes  map[int]*Bike
	nextID int
	mu     sync.Mutex
}

func NewRepository() *Repository {
	return &Repository{
		bikes:  make(map[int]*Bike),
		nextID: 1,
	}
}

func (r *Repository) FindBike(id int) (*Bike, error) {
	bike, exists := r.bikes[id]
	if !exists {
		return nil, fmt.Errorf("Bike not found.")
	}
	return bike, nil
}

func (r *Repository) SaveBike(bike *Bike) {
	r.mu.Lock()
	defer r.mu.Unlock()
	bike.ID = r.nextID
	r.bikes[bike.ID] = bike
	r.nextID++
}

func (r *Repository) PrintBikes() {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, bike := range r.bikes {
		fmt.Println(bike)
	}
}
