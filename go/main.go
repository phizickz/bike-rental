package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"bike-rental/internal/bike"
	"bike-rental/internal/config"
	"bike-rental/internal/customer"
	"bike-rental/internal/rental"
)

func main() {
	config := config.LoadConfig()

	// Initialize database connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)
	db = database.db(dsn)
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// }

	bikeRepo := bike.NewRepository(db)
	customerRepo := customer.NewRepository(db)
	rentalRepo := rental.NewRepository(db)
	// rentalService := rental.NewService(bikeRepo, customerRepo, rentalRepo)
	// reader := bufio.NewReader(os.Stdin)
	SeedData(bikeRepo, customerRepo)
	fmt.Println("Printing bikes")
	fmt.Print(bikeRepo.GetAllBikes())
	fmt.Println("Printing customers")
	fmt.Print(customerRepo.GetAllCustomers())
	fmt.Println("Printing rentals")
	fmt.Print(rentalRepo.GetAllRentals())
	// for {
	// 	PrintOptions()
	// 	input := GetUserInput(reader, "Choose an option: ")

	// 	switch input {
	// 	case "1":
	// 		rentBike(reader, rentalService)
	// 	case "2":
	// 		returnBike(reader, rentalService)
	// 	case "3":
	// 		registerCustomer(reader, customerRepo)
	// 	case "4":
	// 		registerBike(reader, bikeRepo)
	// 	case "5":
	// 		fmt.Println("Goodbye.")
	// 		return
	// 	case "9":
	// 		bikeRepo.PrintBikes()
	// 		customerRepo.PrintCustomers()
	// 		rentalService.PrintRentals()
	// 	default:
	// 		fmt.Println("Invalid option. Try again.")
	// 	}
	// }
}

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

func returnBike(reader *bufio.Reader, rentalService *rental.Service) {
	inputRentalID, err := strconv.Atoi(GetUserInput(reader, "Enter rental ID: "))
	if err != nil {
		fmt.Errorf("Invalid rental ID.")
		return
	}

	rentalService.ReturnBike(inputRentalID)
}

func registerCustomer(reader *bufio.Reader, customerRepo *customer.Repository) {
	input := GetUserInput(reader, "New customer name: ")
	customerRepo.SaveCustomer(customer.NewCustomer(input))
	fmt.Printf("User created.\n")
}

func registerBike(reader *bufio.Reader, bikeRepo *bike.Repository) {
	input := GetUserInput(reader, "New bike model: ")
	bikeRepo.SaveBike(bike.NewBike(input))
}

func PrintOptions() {
	fmt.Println("\nWelcome to the Bike Rental Shop.")
	fmt.Println("1. Rent a bike")
	fmt.Println("2. Return a bike")
	fmt.Println("3. Register a customer")
	fmt.Println("4. Register a bike")
	fmt.Println("5. Exit")
	fmt.Println("9. Infodump")
}

func GetUserInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func SeedData(bikeRepo *bike.Repository, customerRepo *customer.Repository) {
	bikeRepo.SaveBike(bike.NewBike("DBS"))
	bikeRepo.SaveBike(bike.NewBike("Merida"))
	bikeRepo.SaveBike(bike.NewBike("Zykkel"))

	customerRepo.SaveCustomer(customer.NewCustomer("Markus"))
	customerRepo.SaveCustomer(customer.NewCustomer("Marte"))
}
