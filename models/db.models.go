package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User => Estructura usuario
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string             `bson:"username" json:"username,omitempty"`
	Password   string             `bson:"password,omitempty" json:"password,omitempty"`
	CreatedAt  time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	Rol        string             `bson:"rol" json:"rol,omitempty"`
	Nombre     string             `bson:"nombre" json:"nombre,omitempty"`
	Apellido   string             `bson:"apellido" json:"apellido,omitempty"`
	DNI        string             `bson:"dni" json:"dni,omitempty"`
	Telefono   string             `bson:"telefono" json:"telefono,omitempty"`
	Email      string             `bson:"email" json:"email,omitempty"`
	LoginCount int                `bson:"loginCount,omitempty" json:"loginCount,omitempty"`
}
