# Name, age
## Rentalhistory

class Customer:
    def __init__(self, name, phonenumber=0, age=0):
        self._name = name
        self._age = age
        self._phonenumber = phonenumber
        self._id = -1


    @property
    def name(self):
        return self._name

    @name.setter
    def name(self, value):
        self._name = value

    @property
    def age(self):
        return self._age

    @age.setter
    def age(self, value):
        self._age = value

    @property
    def phonenumber(self):
        return self._phonenumber

    @phonenumber.setter
    def phonenumber(self, value):
        self._phonenumber = value

    @property
    def id(self):
        return self._id

    @id.setter
    def id(self, value):
        if value < 0:
            return
        self._id = value