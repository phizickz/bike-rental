# Bikes
class Inventory:
    def __init__(self):
        self.__bikes = []

    def addOneBike(self, bike):
        if bike in self.__bikes:
            raise Exception("Bike already exists in inventory.")

        self.__bikes.append(bike)

    def getOneBike(self, bikeSN):
        pass

    def getAllBikes(self):
        return self.__bikes