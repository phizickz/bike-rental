package customer

import "fmt"

type Customer struct {
	ID   int
	Name string
}

func NewCustomer(name string) *Customer {
	return &Customer{
		Name: name,
	}
}

func (c *Customer) String() string {
	return fmt.Sprintf("Customer ID: %d, name: %s", c.ID, c.Name)
}
