package rental

import (
	"fmt"
	"sync"

	"bike-rental/internal/bike"
	"bike-rental/internal/customer"
)

type Service struct {
	bikeRepo     *bike.Repository
	customerRepo *customer.Repository
	rentals      map[int]*Rental
	nextID       int
	mu           sync.Mutex
}

func NewService(bikeRepo *bike.Repository, customerRepo *customer.Repository) *Service {
	return &Service{
		bikeRepo:     bikeRepo,
		customerRepo: customerRepo,
		rentals:      make(map[int]*Rental),
		nextID:       1,
	}
}

func (s *Service) RentBike(customerID, bikeID int) (*Rental, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	bike, err := s.bikeRepo.FindBike(bikeID)
	if err != nil {
		return nil, err
	}
	if !bike.Available {
		return nil, fmt.Errorf("Bike is not available.")
	}
	id := s.nextID
	rental := NewRental(id, bikeID, customerID)
	s.rentals[rental.ID] = rental
	s.nextID++
	bike.Available = false
	s.bikeRepo.SaveBike(bike)

	// Additional rental logic here (e.g., save rental info, update customer info)

	return rental, nil
}

func (s *Service) ReturnBike(rentalID int) error {
	rental, exists := s.rentals[rentalID]
	if !exists {
		return fmt.Errorf("Rental not found.")
	}
	if rental.IsCompleted() {
		return fmt.Errorf("Rental already completed.")
	}

	rental.CompleteRental()

	bike, err := s.bikeRepo.FindBike(rental.BikeID)
	if err != nil {
		return err
	}
	bike.Available = true
	s.bikeRepo.SaveBike(bike)

	return nil
}

func (s *Service) PrintRentals() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, rental := range s.rentals {
		fmt.Println(rental)
	}
}
