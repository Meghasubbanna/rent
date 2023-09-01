package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
type Book struct {
	//ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Book_id          string `json:"book_id" bson:"book_id"`
	Book_name        string `json:"book_name" bson:"book_name"`
	Rentee_id        string `json:"rentee_id" bson:"rentee_id"`
	Book_Description string `json:"book_description" bson:"book_description"`
	Book_Price       int    `json:"book_price" bson:"book_price"`
	Book_Author      string `json:"book_author" bson:"book_author"`
	Book_Category    string `json:"book_category" bson:"book_category"`
}

type BuyRentBook struct {
	Renter_id          string    `json:"renter_id" bson:"renter_id"`
	Renter_name        string    `json:"renter_name" bson:"renter_name"`
	Book_name          string    `json:"book_name" bson:"book_name"`
	Book_Category      string    `json:"book_category" bson:"book_category"`
	Book_Price         int       `json:"book_price" bson:"book_price"`
	Rentee_id          string    `json:"rentee_id" bson:"rentee_id"`
	Book_Purchase_time time.Time `json:"book_purchase_time" bson:"book_purchase_time"`
	Book_Return_time   time.Time `json:"book_return_time" bson:"book_return_time"`
}
