package main

import (
	"log"

	"Test-api-mongodb/db"
	"Test-api-mongodb/handlers"
)

func main() {

	if !db.TestConnection() {
		log.Fatal("Sin conexion")
		return
	}
	handlers.Handlers()
}
