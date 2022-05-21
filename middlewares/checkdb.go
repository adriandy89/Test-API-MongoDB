package middlewares

import (
	"net/http"

	"Test-api-mongodb/db"
)

// CheckDB => Verifica la conexion a la base de datos
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.TestConnection() {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
