package models

// Response x
type Response struct {
	Message string `bson:"message" json:"message,omitempty"`
}

// Response Error
type ResponseError struct {
	Error string `bson:"error" json:"error,omitempty"`
}
