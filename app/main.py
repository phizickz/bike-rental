import os
import random

from modules.bike import Bike
from modules.customer import Customer
from modules.rental import Rental
from random import choice
from dotenv import load_dotenv
from modules.postgresclient import pgclient

def generateBikes(numberOfBikes: int):
    models = {'Kingdom': 'KNG', 'Sprinter': 'SPR', 'Trailster': 'TRS', 'Flamingo': 'FLM', 'Phoenix': 'PHX', 'Firestorm': 'FST', 'Viper': 'VIP', 'Cobblestone': 'CBL', 'Thunder': 'THR', 'Sniper': 'SNP'}

    year = []
    prices = []
    for y in range(1980, 2020, 2):
        year.append(y)
    for p in range(100, 1000, 50):
        prices.append(p)
    for b in range(numberOfBikes):
        tempModel = choice(list(models.keys()))
        tempYear = choice(year)
        tempBike = Bike(
            model=tempModel,
            price=choice(prices),
            year=tempYear,
            sn=f'{models[tempModel]}-{tempYear}-{b}'
        )
        client.addBike(tempBike)

def generateCustomers(numberOfCustomers: int):
    firstnames = ["Emma", "Lars", "Ingrid", "Oscar", "Ida", "Henrik", "Nora", "Erik", "Astrid", "Magnus"]
    lastnames = ["Søderlund", "Krogh", "Nygård", "Lien", "Fjeld", "Eriksen", "Bjerke", "Haugen", "Bakken", "Myhre"]
    ages = []
    for y in range(18, 50, 2):
        ages.append(y)

    for c in range(numberOfCustomers):
        numberBase = 22345678
        tempphonenumber = numberBase + c if ((numberBase + c) % 1) == 0 else 0
        tempCustomer = Customer(
            name=f"{choice(firstnames)} {choice(lastnames)}",
            age=choice(ages),
            phonenumber=tempphonenumber
        )
        client.addCustomer(tempCustomer)

def generateRentals(numberofRentals: int):
    for r in range(numberofRentals):
        randbike = client.getRandomBikeID(random.randint(1,3))
        randcustomer = client.getRandomCustomerID()
        for b in randbike:
            client.rentBike(
                Rental(
                    bikeid=b[0],
                    customerid=randcustomer[0][0]
                )
            )

if __name__ == '__main__':
    load_dotenv()
    client = pgclient()
    if os.getenv("PY_ENV") == "dev":
        generateBikes(100)
        generateCustomers(50)
        generateRentals(10)
        rentals = client.getRentals()
        for r in rentals:
            print(r)
        # create 20 fictional rentals
        # end half of rentals
        # list out current rentals
        # list out rentalhistory

