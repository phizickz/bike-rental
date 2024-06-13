package rental

import "time"

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
