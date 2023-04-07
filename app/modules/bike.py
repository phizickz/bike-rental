# Model, year, price
## Rentalstate

class Bike:
    def __init__(self, model, sn, year, price, id):
        self.__model: str = model
        self.__year: int = year
        self.__price = price
        self.__sn: str = sn
        self.__isRented: bool = False
        self.__id = id

    def __str__(self):
        return f'Model: {self.__model}\nYear: {self.__year}\nPrice: {self.__price}\nSN: {self.__sn}\nRental state: {self.__isRented}'

    def getModel(self) -> str:
        return self.__model

    def getYear(self) -> int:
        return self.__year

    def getPrice(self) -> int:
        return self.__price

    def setPrice(self, price):
        if price <= 0:
            raise Exception("Price of bike cannot be negative.")
        self.__price = price

    def getSN(self) -> str:
        return self.__dn

    def getIsRented(self) -> bool:
        return self.__isRented

    def setIsRented(self, state: bool):
        self.__isRented = state

    def getID(self) -> int:
        return self.__id
