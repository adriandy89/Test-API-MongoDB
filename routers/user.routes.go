package routers

import (
	"Test-api-mongodb/controllers/usercontroller"
	"Test-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// UserRoutes => Rutas de Usuarios
func UserRoutes(r *mux.Router) {
	//	r.HandleFunc("/api/user/new", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.UserRegister))).Methods("POST")
	r.HandleFunc("/api/user/new", middlewares.CheckDB(usercontroller.UserRegister)).Methods("POST") // Sin token para que puedan crearse su propio user
	r.HandleFunc("/api/user/{id}", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.GetUserByID))).Methods("GET")
	r.HandleFunc("/api/users", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.GetAllUsers))).Methods("GET")
	r.HandleFunc("/api/user/{id}", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.DeleteUserByID))).Methods("DELETE")
	r.HandleFunc("/api/user/{id}", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.UpdateUserByID))).Methods("PUT")
}
