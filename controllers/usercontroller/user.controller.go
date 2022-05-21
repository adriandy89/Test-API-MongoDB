package usercontroller

import (
	"encoding/json"
	"net/http"

	"Test-api-mongodb/models"
	"Test-api-mongodb/services/errorservice"
	"Test-api-mongodb/services/messageservice"
	"Test-api-mongodb/services/userservice"

	"github.com/gorilla/mux"
)

// UserRegister => controlador de la ruta de registro de usuario
func UserRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var userFounded bool = userservice.ValidateIfUserExistByUsername(user.Username)

		if userFounded {
			errorservice.ErrorMessage(w, "Ese usuario ya existe", 400)
			return
		}

		errr := userservice.InsertNewUser(user)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos", 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Usuario creado correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// GetUserByID => obtener un solo usuario mediante un id en los parametros
func GetUserByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var idUser string = vars["id"]
		if len(idUser) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		userFounded, founded := userservice.FindByID(idUser)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		userResponse, err := json.Marshal(userFounded)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion interna de datos", 500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(userResponse)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// GetAllUsers => function to get all enable users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "SA" {

		userList, founded := userservice.FindAll()
		if !founded {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(userList)

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeleteUserByID => eliminar un solo usuario mediante un id en los parametros
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var idUser string = vars["id"]
		if len(idUser) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := userservice.DeleteByID(idUser)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Usuario eliminado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// UpdateUserByID => actualiza un usuario mediante un id en los parametros
func UpdateUserByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var idUser string = vars["id"]
		if len(idUser) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if user.Nombre == "" || user.Username == "" || user.Rol == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		userUpdate, ext := userservice.FindByID(idUser)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		var userFounded bool = userservice.ValidateIfUserExistByUsername(user.Username)

		if userFounded && userUpdate.Username != user.Username {
			errorservice.ErrorMessage(w, "Ese usuario ya está registrado en la base de datos", 400)
			return
		}
		count, err := userservice.UpdateByID(idUser, user)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Usuario actualizado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
