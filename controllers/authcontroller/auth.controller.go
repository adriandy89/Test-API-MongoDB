package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"Test-api-mongodb/models"
	authservice "Test-api-mongodb/services/authservices"
	"Test-api-mongodb/services/errorservice"
)

// Login => Ruta de autenticacion
func Login(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorservice.ErrorMessage(w, "Usuario y/o contraseña invalidos", 400)
		return
	}

	userLogged, exist := authservice.Login(user.Username, user.Password)
	if !exist {
		errorservice.ErrorMessage(w, "Credenciales Invalidas ó Inactivo!", 403)
		return
	}

	userLogged.Password = ""
	userLogged.LoginCount++

	jwtKey, err := generarJwt(userLogged)
	if err != nil {
		errorservice.ErrorMessage(w, "Error en la generaciond de token de autenticación", 500)
		return
	}

	resp := models.LoginReponse{
		Token: jwtKey,
		User:  userLogged,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}

func generarJwt(user models.User) (string, error) {

	myKey := []byte("go-mongo-apu-key-to-secret")
	payload := jwt.MapClaims{
		"username": user.Username,
		"rol":      user.Rol,
		"id":       user.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 6).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
