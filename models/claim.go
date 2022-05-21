package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Claim struct
type Claim struct {
	Username string             `json:"username"`
	Rol      string             `json:"rol"`
	ID       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
