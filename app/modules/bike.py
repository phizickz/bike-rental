# Model, year, price
## Rentalstate

class Bike:
    def __init__(self, model, sn, year, price):
        self._model: str = model
        self._year: int = year
        self._price = price
        self._serialnr: str = sn
        self._id = -1

    @property
    def model(self):
        return self._model

    @property
    def year(self):
        return self._year

    @property
    def price(self):
        return self._price

    @price.setter
    def price(self, value):
        if value <= 0:
            raise Exception("Price of bike cannot be negative.")
        self._price = value

    @property
    def sn(self):
        return self._serialnr


    @property
    def id(self):
        return self._id

    @id.setter
    def id(self, value):
        if value < 0:
            Exception("Bike ID cannot be less than 0")
        self._id = value
