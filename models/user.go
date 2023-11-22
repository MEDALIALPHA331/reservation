package models

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	//! should not be above
	bcryptCost = 5
)

type User struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"omitempty"`
	FirstName string `bson:"firstName,omitempty" json:"firstName"`
	LastName string `bson:"lastName" json:"lastName"`
	Email    string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"EncryptedPassword" json:"-"`
}

type UserDTO struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"Password"`
}


//? Warning This function can be a bottleneck for perf!!
func CreateUserFromDTO(user *UserDTO) *User {
	encpw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcryptCost)

	if err != nil {
		log.Fatalf("Could not encrypt password, %+v", err)
	}

	return &User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email:    user.Email,
		EncryptedPassword: string(encpw),
	}
}
