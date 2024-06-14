package rental

import (
	"fmt"
	"time"
)

type Rental struct {
	ID         int
	BikeID     int
	CustomerID int
	StartTime  time.Time
	EndTime    time.Time
}

func NewRental(id int, bikeID int, customerID int) *Rental {
	return &Rental{
		ID:         id,
		BikeID:     bikeID,
		CustomerID: customerID,
		StartTime:  time.Now(),
	}
}

func (r *Rental) CompleteRental() {
	r.EndTime = time.Now()
}

func (r *Rental) IsCompleted() bool {
	return !r.EndTime.IsZero()
}

func (r *Rental) String() string {
	return fmt.Sprintf("Rental ID: %d, bike ID: %d, customer ID: %d, Start time: %s", r.ID, r.BikeID, r.CustomerID, r.StartTime)
}
