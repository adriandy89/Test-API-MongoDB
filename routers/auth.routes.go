package routers

import (
	"Test-api-mongodb/controllers/authcontroller"
	"Test-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// AuthRoutes => Rutas de validacion y autenticacion
func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/api/auth/login", middlewares.CheckDB(authcontroller.Login)).Methods("POST")
}
