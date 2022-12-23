package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// defining models

// User Model
type User struct{
	Id primitive.ObjectID `bson:"_id,omitempty"`
	FullName string	`bson:"fullname,omitempty"`
	Email string	`bson:"email,omitempty"`
	Password string	`bson:"password,-"`
	DOB string	`bson:"dob,omitempty"`
	Gender string	`bson:"gender,omitempty"`
}

