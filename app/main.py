import os

from modules.bike import Bike
from modules.customer import Customer
from random import choice
from dotenv import load_dotenv
from modules.postgresclient import pgclient

def fillInventory(numberOfBikes: int):
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
            sn=f'{models[tempModel]}-{tempYear}'
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

if __name__ == '__main__':
    load_dotenv()
    client = pgclient()
    if os.getenv("PY_ENV") == "dev":
        fillInventory(500)
        generateCustomers(200)

        # print(f"{client.getBikeAmount()} bikes and {client.getCustomerAmount()} customers in dev-database.")
