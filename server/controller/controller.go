package controller

import (
	"context"
	"encoding/json"

	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/Vansh0100/E-Library/model"
	// "github.com/Vansh0100/E-Library/model"
	"github.com/Vansh0100/E-Library/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database
var userCollection *mongo.Collection

func Connection(mongoUrl string) {
	newClient := options.Client().ApplyURI(mongoUrl)

	client, err := mongo.Connect(context.TODO(), newClient)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo DB Connection successfull!")
	Controller(client)
}
func Controller(client *mongo.Client) {
	database = client.Database("E-Library")
	userCollection = database.Collection("User")
	// data:=model.User{FullName: "Vansh Jaiswal",Email: "Vansh238@gmail.com",Password: "12345",DOB: "20/01/2000",Gender: "Male"}
	// insertedId:=registerNewUser(&data)
	// fmt.Println("Inserted Id:",insertedId)
	// data:=getAllUsers()
	// finalJson,_:=json.MarshalIndent(data,"","\t")
	// fmt.Println(string(finalJson))
	// validateLogin("Vansh23@gmail.com")
	// fmt.Println(data.Email)
}



// HomePage
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to E-Library!</h1>"))
}
// functions to handle data with mongo db

func validateLogin(emailid string) model.User {
	var result model.User
	filter:=bson.D{{"email",emailid}}
	data:=userCollection.FindOne(context.Background(),filter).Decode(&result)
	if data!=nil{
		return model.User{}
	}
	
	return result
}

// Verify Login Credentials
func VerifyLogin(w http.ResponseWriter,r *http.Request)  {
	var enteredCredentials model.User
	json.NewDecoder(r.Body).Decode(&enteredCredentials)
	email:=enteredCredentials.Email
	password:=enteredCredentials.Password

	data:=validateLogin(email)
	// json.NewEncoder(w).Encode(data)
	
	if data.Email==email{
		if data.Password==password{
	json.NewEncoder(w).Encode("Login Successfull")

		}else {
	json.NewEncoder(w).Encode("Wrong Password")
			
		}
	}else {
	json.NewEncoder(w).Encode("User doesn't exists!")
		
	}
}

func registerNewUser(data model.User) interface{} {
	
	insert,err:=userCollection.InsertOne(context.Background(),data)
	if err!=nil{
		return nil
	}else {
		return insert.InsertedID
	}
}

func RegisterUser(w http.ResponseWriter,r *http.Request)  {
	var result model.User
	json.NewDecoder(r.Body).Decode(&result)
	if !checkEmptyFields(&result){
	json.NewEncoder(w).Encode("Necessary credentials are empty!")
		
	}else {
		isExistingUser:=validateLogin(result.Email)
		if isExistingUser.Email==result.Email{
	json.NewEncoder(w).Encode("User already exists! Please Login")

		}else {
			insertedId:=registerNewUser(result)
			if insertedId==nil{
	json.NewEncoder(w).Encode("Data not inserted successfully!")

			}else {
				
	json.NewEncoder(w).Encode("Inserted ID:")
	json.NewEncoder(w).Encode(insertedId)
				
			}
		}
	}
}

func checkEmptyFields(data *model.User) bool {
	return len(data.FullName)!=0 && len(data.Email)!=0 && len(data.Password)!=0
}