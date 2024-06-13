package rental

import (
	"fmt"
	"go/internal/bike"
	"go/internal/customer"
)

type Service struct {
	bikeRepo     *bike.Repository
	customerRepo *customer.Repository
	rentals      map[int]*Rental
	nextID       int
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
	bike, err := s.bikeRepo.FindBike(bikeID)
	if err != nil {
		return nil, err
	}
	if !bike.Available {
		return nil, fmt.Errorf("Bike is not available.")
	}

	rental := NewRental(s.IncrementID(), bikeID, customerID)
	s.rentals[rental.ID] = rental
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

func (s *Service) IncrementID() int {
	id := s.nextID
	s.nextID++
	return id
}
