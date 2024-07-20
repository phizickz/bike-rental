package rental

import (
	"fmt"
	"time"

	"bike-rental/internal/bike"
	"bike-rental/internal/customer"
)

type Service struct {
	bikeRepo     *bike.Repository
	customerRepo *customer.Repository
	rentalRepo   *Repository
}

func NewService(bikeRepo *bike.Repository, customerRepo *customer.Repository, rentalRepo *Repository) *Service {
	return &Service{
		bikeRepo:     bikeRepo,
		customerRepo: customerRepo,
		rentalRepo:   rentalRepo,
	}
}

func (s *Service) RentBike(customerID, bikeID int) (*Rental, error) {
	bike, err := s.bikeRepo.FindBike(bikeID)
	if err != nil {
		return nil, err
	}
	if !bike.Available {
		return nil, fmt.Errorf("Bike is not available")
	}

	rental := &Rental{
		BikeID:     bikeID,
		CustomerID: customerID,
		StartTime:  time.Now(),
	}
	if err := s.rentalRepo.CreateRental(rental); err != nil {
		return nil, err
	}

	// Update bike availability
	bike.Available = false
	if err := s.bikeRepo.SaveBike(bike); err != nil {
		return nil, err
	}

	return rental, nil
}

func (s *Service) ReturnBike(rentalID int) error {
	rental, err := s.rentalRepo.FindRental(rentalID)
	if err != nil {
		return err
	}

	if !rental.EndTime.IsZero() {
		return fmt.Errorf("Rental has been ended")
	}

	rental.EndTime = time.Now()
	if err := s.rentalRepo.SaveRental(rental); err != nil {
		return err
	}
	// Update bike availability
	bike, err := s.bikeRepo.FindBike(rental.BikeID)
	if err != nil {
		return err
	}
	bike.Available = true
	if err := s.bikeRepo.SaveBike(bike); err != nil {
		return err
	}
	return nil
}
