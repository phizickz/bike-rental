package customer

type Customer struct {
	ID   int
	Name string
}

func NewCustomer(id int, name string) *Customer {
	return &Customer{
		ID:   id,
		Name: name,
	}
}
