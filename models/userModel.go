package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID
	FirstName string
	LastName  string
	Email     string
	PhoneNo   string
	
}
