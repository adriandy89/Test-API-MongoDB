package models

// LoginReponse => Cuerpo de respuesta la momento de hacer login
type LoginReponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
