package controller

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	// _ "time"

	// "github.com/omprakas123/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/omprakas123/database"
	"github.com/omprakas123/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type RenteeInfo struct {
	Rentee_id string `json:"rentee_id" bson:"rentee_id"`
}

var RenteeUser RenteeInfo

var Secret_key = []byte("Cochin university")

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func Validate(user models.User, w http.ResponseWriter) bool {
	if user.User_type == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"User Id cannot be empty"}`))
		return false
	} else if user.First_name == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"First Name cannot be empty"}`))
		return false
	} else if user.Last_name == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Last name must be Enter"}`))
		return false
	} else if len(user.Password) < 5 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Password length should be greater than 5"}`))
		return false
	}
	return true
}

func AddRenter(user models.User, w http.ResponseWriter) {
	var rentalUser models.Renter
	rentalUser.User_id = user.User_id
	rentalUser.Name = user.First_name + " " + user.Last_name
	rentalUser.Rental_Email = user.Email

	RentalCol := database.Client1.Database("Rentaluser").Collection("Rental")

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	res, _ := RentalCol.InsertOne(ctx, rentalUser)
	json.NewEncoder(w).Encode(res)

}
func GenerateOwnJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(Secret_key)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

func AddRentee(user models.User, w http.ResponseWriter) {
	var renteeUser models.Rentee
	renteeUser.User_id = user.User_id
	renteeUser.Name = user.First_name + " " + user.Last_name
	renteeUser.Rentee_Email = user.Email
	RenteeUser.Rentee_id = user.User_id
	RenteeCol := database.Client1.Database("Rentaluser").Collection("Rentee")
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	res, _ := RenteeCol.InsertOne(ctx, renteeUser)
	json.NewEncoder(w).Encode(res)
	fmt.Println("Hey! the user is rentee so you can add a book")
}
func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	//var dbUser model.User
	json.NewDecoder(r.Body).Decode(&user)

	if !Validate(user, w) {
		return
	}
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	user.Password = getHash([]byte(user.Password))
	//   fmt.Println(user)
	collection := database.Client1.Database("Rentaluser").Collection("userdata")
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	result, err := collection.InsertOne(ctx, user)

	fmt.Println(err)
	json.NewEncoder(w).Encode(result)

	if strings.ToLower(user.User_type) == "rental" {
		AddRenter(user, w)
	} else if strings.ToLower(user.User_type) == "rentee" {
		AddRentee(user, w)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	var checkdb models.User
	json.NewDecoder(r.Body).Decode(&user)

	collection := database.Client1.Database("Rentaluser").Collection("userdata")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&checkdb)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	userPass := []byte(user.Password)
	CheckPass := []byte(checkdb.Password)
	passErrCheck := bcrypt.CompareHashAndPassword(CheckPass, userPass)

	if passErrCheck != nil {
		log.Println(passErrCheck)
		w.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}

	jwtToken, err := GenerateOwnJWT()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	w.Write([]byte(`{"token":"` + jwtToken + `"}`))
}

func BookCreation(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("hello")
	w.Header().Set("Content-Type", "application/json")
	var user models.Book
	json.NewDecoder(r.Body).Decode(&user)
	user.Rentee_id = RenteeUser.Rentee_id
	user.Book_id = primitive.NewObjectID().Hex()
	fmt.Println(user.Book_id)
	// fmt.Println(user)
	collection1 := database.Client1.Database("Rentaluser").Collection("Book")
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	result, err := collection1.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
	RenteeUser = RenteeInfo{}
}

func AvailableBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	BookAvailability := r.URL.Query().Get("bookcategory")
	//fmt.Println(BookAvailability)
	BookCollection, err := GetBookCategory(BookAvailability)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"This Book category is not available in the database"}`))
		return
	}
	json.NewEncoder(w).Encode(BookCollection)
}

func GetBookCategory(BookAvailability string) ([]models.Book, error) {
	collection1 := database.Client1.Database("Rentaluser").Collection("Book")
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	filterBook := bson.M{"book_category": BookAvailability}
	Bookstore, err := collection1.Find(ctx, filterBook)
	if err != nil {
		return nil, err
	}
	defer Bookstore.Close(ctx)
	var Bookcategory []models.Book
	for Bookstore.Next(ctx) {
		var bookDetail models.Book
		if err := Bookstore.Decode(&bookDetail); err != nil {
			return nil, err
		}
		fmt.Println(Bookstore)
		Bookcategory = append(Bookcategory, bookDetail)
	}
	if err := Bookstore.Err(); err != nil {
		return nil, err
	}
	//fmt.Println(Bookcategory)
	return Bookcategory, nil
}

func BookPurchase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	collection1 := database.Client1.Database("Rentaluser").Collection("Book")
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give a book name:")
	bookname, _ := reader.ReadString('\n')
	fmt.Println(bookname)
	filterBook := bson.M{"Book_name": bookname}
	var result models.Book
	project := bson.D{{"Book_name", 1}}
	opts := options.FindOne().SetProjection(project)
	if err := collection1.FindOne(ctx, filterBook, opts).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Book not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Fatal(err)
		}
		return

	}

	// err := collection1.FindOne(ctx, filterBook, opts).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// collection2 := database.Client1.Database("Rentaluser").Collection("BuyRentBook")
	// var userdata models.BuyRentBook
	// json.NewDecoder(r.Body).Decode(&userdata)
	// userdata.Rentee_id = GetaBook.Rentee_id
	json.NewEncoder(w).Encode(result)

	//fmt.Println(result)
}
