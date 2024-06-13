package bike

import "fmt"

type Repository struct {
	bikes map[int]*Bike
}

func NewRepository() *Repository {
	return &Repository{
		bikes: make(map[int]*Bike),
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
	r.bikes[bike.ID] = bike
}

func (r *Repository) GetHighestID() int {
	return len(r.bikes)
}
