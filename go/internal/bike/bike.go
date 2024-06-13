package bike

type Bike struct {
	ID int
	Model string
	Available bool
}

func NewBike(id int, model string) *Bike {
	return &Bike{
		ID: id,
		Model: model,
		Available: true,
	}
}