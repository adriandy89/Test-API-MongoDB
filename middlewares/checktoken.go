package middlewares

import (
	"errors"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"Test-api-mongodb/models"
	"Test-api-mongodb/services/errorservice"
	"Test-api-mongodb/services/userservice"
)

// UserLogged => Usuario que ha hecho en la operación
var UserLogged models.User

// ValidateJWT func
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		founded, rol, err := validateToken(r.Header.Get("Authorization"))
		if err != nil {
			errorservice.ErrorMessage(w, "Error en validacion de autorización", 401)
			return
		}
		if !founded {
			errorservice.ErrorMessage(w, "El usuario no existe o Token Invalido", 401)
			return
		}
		r.Header.Add("rol", rol)

		next.ServeHTTP(w, r)

	}
}

func validateToken(token string) (bool, string, error) {
	myKey := []byte("go-mongo-apu-key-to-secret")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer ")
	if len(splitToken) != 2 {
		return false, string(""), errors.New("token invalido")
	}

	cleanToken := strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(cleanToken, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil && claims.Username == "SA" {
		return true, claims.Rol, nil
	} else if err == nil && claims.Username != "SA" {
		founded := userservice.ValidateIfUserExistByUsername(claims.Username)
		return founded, claims.Rol, nil
	}

	if !tkn.Valid {
		return false, string(""), errors.New("token invalido")
	}

	return false, claims.Rol, err

}
