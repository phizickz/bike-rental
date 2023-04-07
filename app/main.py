from modules.inventory import Inventory
from modules.bike import Bike
from random import choice
def fillInventory(numberOfBikes: int):
    models = {
        "Merida": "mi",
        "DBS": "db",
        "Samsung": "ss"
    }
    year = [2000, 2010, 2020]
    prices = [20, 30, 40]
    for b in range(numberOfBikes):
        tempModel = choice(list(models.keys()))
        tempYear = choice(year)
        inventory.addOneBike(
            Bike(
                model=tempModel,
                price=choice(prices),
                year=tempYear,
                sn=f'{models[tempModel]}-{tempYear}',
                id = b + 1
            )
        )
    print("Inventory filled with bikes.")

if __name__ == '__main__':
    inventory = Inventory()
    fillInventory(30)
