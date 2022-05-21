package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	MongoURI    = goDotEnvVariable("mongoURI")
	MongoDBName = goDotEnvVariable("mongoDBName")
	ServerPORT  = goDotEnvVariable("serverPORT")
	FQDN        = goDotEnvVariable("FQDN")
	Domain      = goDotEnvVariable("domain")
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
