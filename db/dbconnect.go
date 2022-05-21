package db

import (
	"context"
	"log"

	config "Test-api-mongodb/config_loader"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN => Se exporta conexion a mongo para utilizar en toda la app
var MongoCN = conectarDB()

var clientOptions = options.Client().ApplyURI(config.MongoURI)

// DB => db name
var DB = MongoCN.Database(config.MongoDBName)

// Colllections -------
var UserCollection = DB.Collection("users")
var CategoryCollection = DB.Collection("categories")
var ClientProviderCollection = DB.Collection("clientsProviders")
var ContractCollection = DB.Collection("contracts")
var ContractNonExecutionCollection = DB.Collection("contractsNonExecutions")
var ContractSpecificCollection = DB.Collection("contractsSpecific")
var CountryCollection = DB.Collection("countries")
var CurrencyCollection = DB.Collection("currencies")
var DescriptionsStatesCollection = DB.Collection("descriptionsStates")
var EntityCollection = DB.Collection("entities")
var NonEjecutionsCollection = DB.Collection("nonEjecutions")
var ObjectContractsCollection = DB.Collection("objectContracts")
var OffersRequestsCollection = DB.Collection("offersRequest")
var OrganismsCollection = DB.Collection("organisms")
var SuplementCollection = DB.Collection("supplements")
var SuplementSpecificCollection = DB.Collection("supplementsSpecific")
var SupplementsOperationsCollection = DB.Collection("supplementsOperations")
var TypeContractsCollection = DB.Collection("typeContracts")
var TypeFactsCollection = DB.Collection("typeFacts")
var SectorCollection = DB.Collection("sectors")

func conectarDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa a la db")
	return client
}

// TestConnection => Inicializar los valores de la Conexion a la BD y verifica la conexion a la db mediante un ping
func TestConnection() bool {

	err := MongoCN.Ping(context.TODO(), nil)

	return err == nil
}
