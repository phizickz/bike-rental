package main

import (
	"bufio"
	"fmt"
	"go/internal/bike"
	"go/internal/customer"
	"go/internal/rental"
	"os"
	"strconv"
	"strings"
)

func main() {
	// config := config.LoadConfig()

	bikeRepo := bike.NewRepository()
	customerRepo := customer.NewRepository()
	rentalService := rental.NewService(bikeRepo, customerRepo)
	reader := bufio.NewReader(os.Stdin)

	for {
		PrintOptions()
		input := GetUserInput(reader, "Choose an option: ")

		switch input {
		case "1":
			rentBike(reader, rentalService)
		case "2":
			returnBike(reader, rentalService)
		case "3":
			registerCustomer(reader, customerRepo)
		case "4":
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

// TODO add rentalrepository
func rentBike(reader *bufio.Reader, rentalService *rental.Service) {
	customerID, err := strconv.Atoi(GetUserInput(reader, "Enter customer ID: "))
	if err != nil {
		fmt.Println("Invalid customer ID.")
		return
	}

	bikeID, err := strconv.Atoi(GetUserInput(reader, "Enter bike ID: "))
	if err != nil {
		fmt.Println("Invalid bike ID.")
		return
	}
	_, err = rentalService.RentBike(customerID, bikeID)
	if err != nil {
		fmt.Println("Error renting bike: ", err)
		return
	}

	fmt.Println("Bike rented successfully.")
}

// TODO
func returnBike(reader *bufio.Reader, rentalService *rental.Service) {
	inputRentalID := GetUserInput(reader, "Enter rental ID: ")
}

func registerCustomer(reader *bufio.Reader, customerRepo *customer.Repository) {
	input := GetUserInput(reader, "New customer name: ")
	id := customerRepo.GetHighestID
	customerRepo.SaveCustomer(customer.NewCustomer(id, input))
	fmt.Printf("User %v created.\n", id)
}

// TODO
func registerBike() {

}

func PrintOptions() {
	fmt.Println("Welcome to the Bike Rental Shop.")
	fmt.Println("1. Rent a bike")
	fmt.Println("2. Return a bike")
	fmt.Println("3. Register a customer")
	fmt.Println("4. Exit")
}

func GetUserInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func SeedData(bikeRepo *bike.Repository, customerRepo *customer.Repository) {
	bikeRepo.SaveBike(bike.NewBike(bikeRepo.GetHighestID, "DBS"))
	bikeRepo.SaveBike(bike.NewBike(bikeRepo.GetHighestID, "Merida"))
	bikeRepo.SaveBike(bike.NewBike(bikeRepo.GetHighestID, "Zykkel"))

	customerRepo.SaveCustomer(customer.NewCustomer(customerRepo.GetHighestID, "Markus"))
	customerRepo.SaveCustomer(customer.NewCustomer(customerRepo.GetHighestID, "Marte"))
}
