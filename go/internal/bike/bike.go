package bike

import "fmt"

type Bike struct {
	ID        int
	Model     string
	Available bool
}

func NewBike(model string) *Bike {
	return &Bike{
		Model:     model,
		Available: true,
	}
}

func (b *Bike) String() string {
	return fmt.Sprintf("Bike ID: %d, model: %s, available: %t", b.ID, b.Model, b.Available)
}
