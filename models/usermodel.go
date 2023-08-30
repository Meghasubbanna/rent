package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	First_name string             `json:"first_name" bson:"first_name"`
	Last_name  string             `json:"last_name" bson:"last_name"`
	Password   string             `json:"password" bson:"password"`
	Email      string             `json:"email" bson:"email"`
	User_id    string             `json:"user_id" bson:"user_id"`
	User_type  string             `json:"user_type" bson:"usert_type"`
}

type Renter struct {
	User_id      string `json:"Rental_user_id" bson:"Rental_user_id"`
	Name         string `json:"Rental_name" bson:"Rental_name"`
	Rental_Email string `json:"Rental_email" bson:"rental_email"`
}

type Rentee struct {
	User_id      string `json:"Rentee_user_id" bson:"Rentee_user_id"`
	Name         string `json:"Rentee_name" bson:"Rentee_name"`
	Rentee_Email string `json:"Rentee_email" bson:"rentee_email"`
}


