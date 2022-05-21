package userservice

import (
	"context"
	"time"

	"Test-api-mongodb/db"
	"Test-api-mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// ValidateIfUserExistByUsername => buscar un usuario en la base de datos por username
func ValidateIfUserExistByUsername(username string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"username": username}

	var result models.User

	err := db.UserCollection.FindOne(ctx, condition).Decode(&result)

	return err == nil
}

// SaveUser => guarda un nuevo usuario en la base de datos
func InsertNewUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	u.Password, _ = encriptPassword(u.Password)
	u.CreatedAt = time.Now().UTC()

	_, err := db.UserCollection.InsertOne(ctx, u)

	return err
}

// FindByID => Funcion para obtener usuario por id
func FindByID(id string) (models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	userID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": userID}
	var result models.User

	err := db.UserCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	result.Password = ""
	return result, true
}

// FindByUsername => Busca un usuario por el usuario
func FindByUsername(username string) (models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"username": username, "state": 1}

	var userFounded models.User

	err := db.UserCollection.FindOne(ctx, condition).Decode(&userFounded)
	if err != nil {
		return userFounded, false
	}
	return userFounded, true
}

// FindAll => return all users registered by Company
func FindAll() ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var results []*models.User

	cursor, err := db.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return results, false
		}
		user.Password = ""
		results = append(results, &user)
	}

	return results, true
}

// DeleteByID => Funcion para eliminar usuario por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	userID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": userID}

	delRes, err := db.UserCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}

// UpdateByID => Funcion para actualizar usuario por id
func UpdateByID(id string, userUpdate models.User) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	userID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": userID}

	// -----------------------------aki arreglar los datos
	var update primitive.M
	if userUpdate.Password != "" {
		userUpdate.Password, _ = encriptPassword(userUpdate.Password)
		update = bson.M{"$set": bson.M{
			"username":   userUpdate.Username,
			"rol":        userUpdate.Rol,
			"password":   userUpdate.Password,
			"nombre":     userUpdate.Nombre,
			"apellido":   userUpdate.Apellido,
			"dni":        userUpdate.DNI,
			"telefono":   userUpdate.Telefono,
			"email":      userUpdate.Email,
			"loginCount": userUpdate.LoginCount,
		}}
	} else {
		update = bson.M{"$set": bson.M{
			"username":   userUpdate.Username,
			"rol":        userUpdate.Rol,
			"nombre":     userUpdate.Nombre,
			"apellido":   userUpdate.Apellido,
			"dni":        userUpdate.DNI,
			"telefono":   userUpdate.Telefono,
			"email":      userUpdate.Email,
			"loginCount": userUpdate.LoginCount,
		}}
	}
	upd, err := db.UserCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// UpdateLoginCount => Funcion para actualizar usuario por id
func UpdateLoginCount(id primitive.ObjectID, count int) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{
		"loginCount": count + 1,
	}}

	db.UserCollection.UpdateOne(ctx, condition, update)
}

func encriptPassword(password string) (string, error) {
	var cost int = 5
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
