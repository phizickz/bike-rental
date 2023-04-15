import os

import psycopg2
from modules.bike import Bike
from modules.customer import Customer

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

    def getAllTables(self):
        res = self.executeSQL("select relname from pg_class where relkind='r' and relname !~ '^(pg_|sql_)';")
        print(res)

    def addBike(self, value: Bike):
        res = self.executeSQL(f"INSERT INTO bike (model,year,price,serialnr,isRented) VAlUES('{value.model}','{value.year}','{value.price}','{value.sn}','{value.isRented}') RETURNING id;")
        value.id = res[0][0]

    def isPhoneNumberTaken(self, value):
        res = self.executeSQL(f"SELECT * FROM customer WHERE phonenumber={value.phonenumber};")
        if len(res) > 0:
            return True
        return False

    def addCustomer(self, value: Customer):
        if self.isPhoneNumberTaken(value):
            print("Phone number already in use.")
            return

        res = self.executeSQL(f"INSERT INTO customer (name,age,phonenumber) VAlUES('{value.name}','{value.age}','{value.phonenumber}') RETURNING id;")
        value.id = res[0][0]

    def getAllBikes(self):
        return self.executeSQL(f"SELECT * FROM bike")

    def getBikeAmount(self):
        res = self.executeSQL(f"SELECT * FROM bike")
        return len(res)
    def getCustomerAmount(self):
        res = self.executeSQL(f"SELECT * FROM customer")
        return len(res)

    def getAllCustomers(self):
        return self.executeSQL(f"SELECT * FROM customer")

    def executeSQL(self, sqlstring):
        with self.__client as conn:
            with conn.cursor() as curs:
                curs.execute(sqlstring)
                return curs.fetchall()
