import os

import psycopg2
from modules.bike import Bike
from modules.customer import Customer
from modules.rental import Rental

class pgclient:
    def __init__(self):
        try:
            self.__host = os.getenv("DB_HOST")
            self.__port = "5432" if os.getenv("DB_PORT") == "" else os.getenv("DB_PORT")
            self.__database = os.getenv("DB_DATABASE")
            self.__user = os.getenv("DB_USER")
            self.__password = os.getenv("DB_PASSWORD")
            self.__client = psycopg2.connect(
                f"dbname='{self.__database}' user='{self.__user}' host='{self.__host}' port='{self.__port}' password='{self.__password}'"
            )

            self.__client.autocommit = True
        except Exception as e:
            print(e)
            exit(1)

    def __del__(self):
        self.__client.close()


    def addBike(self, value: Bike):
        res = self.executeSQL(f"INSERT INTO bike (model,year,price,serialnr) "
                              f"VAlUES('{value.model}','{value.year}','{value.price}','{value.sn}') "
                              f"RETURNING id;")
        value.id = res[0][0]

    def isPhoneNumberTaken(self, value):
        res = self.executeSQL(f"SELECT * FROM customer WHERE phonenumber={value.phonenumber};")
        if len(res) > 0:
            return True
        return False

    def addCustomer(self, value: Customer):
        if self.isPhoneNumberTaken(value):
            # print("Phone number already in use.")
            return

        res = self.executeSQL(f"INSERT INTO customer (name,age,phonenumber) "
                              f"VAlUES('{value.name}','{value.age}','{value.phonenumber}') "
                              f"RETURNING id;")
        value.id = res[0][0]

    def getRandomBikeID(self, amount=1) -> list:
        # Returns a bike ID that is not currently rented
        # All honour to ChatGPT for query
        res = self.executeSQL(f"SELECT id "
                              f"FROM bike "
                              f"WHERE id NOT IN ("
                              f"SELECT bike_id "
                              f"FROM rentals "
                              f"INNER JOIN activeRentals ON activeRentals.rental_id = rentals.id"
                              f") "
                              f"ORDER BY RANDOM() LIMIT {amount};")
        return res

    def getRandomCustomerID(self, value=1) -> list:
        res = self.executeSQL(f"SELECT id FROM customer ORDER BY RANDOM() LIMIT {value};")
        return res

    def getRandomCustomerIDWithRental(self):
        res = self.executeSQL(f"SELECT rentals.customer_id"
                              f"FROM rentals"
                              f"INNER JOIN activeRentals ON rentals.id = activeRentals.rental_id"
                              f"WHERE customer_id=("
                              f"SELECT customer_id FROM rentals ORDER BY RANDOM() LIMIT 1;"
                              f");")
        return res[0]

    def rentBike(self, value: Rental):
        res = self.executeSQL(f"INSERT INTO rentals(start_date,stop_date,customer_id,bike_id) "
                              f"VALUES('{value.startdate}','{value.stopdate}','{value.customerid}','{value.bikeid}')"
                              f"RETURNING id;")
        if len(res) < 1:
            raise Exception("Something went wrong at rentBike.")
        res = self.executeSQL(f"INSERT INTO activeRentals(rental_id)"
                              f"VALUES('{res[0][0]}')")

    def getRentals(self):
        res = self.executeSQL(f"select * from rentals")
        return res

    def getActiveRentalsForRandomCustomer(self):
        res = self.executeSQL(f"SELECT rentals.customer_id"
                              f"FROM rentals"
                              f"INNER JOIN activeRentals ON rentals.id = activeRentals.rental_id;")
        return res

    def executeSQL(self, sqlstring):
        try:
            with self.__client as conn:
                with conn.cursor() as curs:
                    curs.execute(sqlstring)
                    return curs.fetchall()
        except Exception as e:
            print(e)
    # def getAllTables(self):
    #     res = self.executeSQL("select relname from pg_class where relkind='r' and relname !~ '^(pg_|sql_)';")
    #     print(res)
    # def getAllBikes(self):
    #     return self.executeSQL(f"SELECT * FROM bike")
    #
    # def getBikeAmount(self):
    #     res = self.executeSQL(f"SELECT * FROM bike")
    #     return len(res)
    # def getCustomerAmount(self):
    #     res = self.executeSQL(f"SELECT * FROM customer")
    #     return len(res)
    #
    # def getAllCustomers(self):
    #     return self.executeSQL(f"SELECT * FROM customer")


