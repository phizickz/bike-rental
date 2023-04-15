# Rent bike to customer
##
import datetime

from modules.bike import Bike
from modules.customer import Customer

class Date():
    def __init__(self, date=datetime.__all__):
        self._date = date

    @property
    def date(self):
        return self._date

    @date.setter
    def date(self, value: datetime):
        self._date = value


class Rental():
    def __init__(self,bikeid: int, customerid: int, startdate=Date().date,stopdate=Date()):
        self._bikeid=bikeid
        self._customerid = customerid
        self._startdate = startdate
        self._stopdate = stopdate
        self._id = -1

    @property
    def bikeid(self):
        return self._bikeid

    @property
    def customerid(self):
        return self._customerid

    @property
    def startdate(self):
        return self._startdate

    @property
    def stopdate(self):
        return self._stopdate

    @property
    def id(self):
        return self._id

    @id.setter
    def id(self, value):
        if value < 0:
            return
        self._id = value