from pymongo.mongo_client import MongoClient
from pymongo.server_api import ServerApi
import ssl
class MongodbClient:
    def __init__(self, connectionuri, database):
        self.__client = MongoClient(connectionuri,
                                    tlsAllowInvalidCertificates=True,
                                    server_api=ServerApi('1'))
        self.__db = self.__client[database]
        try:
            self.__client.admin.command('ping')
            print("Pinged your deployment. You successfully connected to MongoDB!")
        except Exception as e:
            print(e)
            exit(1)
    def findDocument(self, doc: dict, coll: str):
        return self.__db[coll].find_one(doc)

    def insertDocument(self, doc: dict, coll: str):
        self.__db[coll].insert_one(doc)

    def insertManyDocuments(self, doc: dict, coll: str):
        self.__db[coll].insert_many(doc)