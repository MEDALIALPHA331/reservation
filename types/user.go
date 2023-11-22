package types


type User struct {
	Id string `bson:"_id,omitempty" json:"omitempty"`
	firstName string `bson:"firstName,omitempty" json:"firstName"`
	lastName string `bson:"lastName" json:"lastName"`
	email string `bson:"email" json:"email"`
	password string `bson:"password" json:"-"`
}