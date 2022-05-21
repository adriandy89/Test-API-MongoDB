package handlers

import (
	config "Test-api-mongodb/config_loader"
	"Test-api-mongodb/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//const FSPATH = "./frontend/"

// Handlers => En la funcion que maneja las peticiones
func Handlers() {

	router := mux.NewRouter()
	routers.AuthRoutes(router)
	routers.UserRoutes(router)

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend"))).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(&withCustom404PageHandler{http.Dir("./frontend")})).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = config.ServerPORT
	}

	log.Println("Servidor Online, Puerto:", config.ServerPORT)

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

type withCustom404PageHandler struct{ fs http.FileSystem }

func (w4h *withCustom404PageHandler) Open(name string) (http.File, error) {
	f, err := w4h.fs.Open(name)
	if os.IsNotExist(err) {
		return w4h.fs.Open("index.html")
	}
	return f, err
}
