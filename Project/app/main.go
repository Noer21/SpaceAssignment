package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type Apartment struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
	Unit int                `json:"unit,omitempty" bson:"unit,omitempty"`
	City string             `json:"city,omitempty" bson:"city,omitempty"`
}

type ReqStatus struct {
	Status string `json:"status,omitempty"`
	Code   int    `json:"code,omitempty"`
}

func CreateApartment(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")

	var apartment Apartment
	json.NewDecoder(req.Body).Decode(&apartment)
	collection := client.Database("StockSpace").Collection("apartments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.InsertOne(ctx, apartment)

	result := ReqStatus{}
	if err == nil {
		result = ReqStatus{
			Status: "Success!",
			Code:   200,
		}
	} else {
		result = ReqStatus{
			Status: "Fail!",
			Code:   500,
		}
	}
	json.NewEncoder(res).Encode(result)
}

func GetAllApartment(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")

	var apartments []Apartment
	collection := client.Database("StockSpace").Collection("apartments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.Find(ctx, Apartment{})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	defer result.Close(ctx)
	for result.Next(ctx) {
		var apartment Apartment
		result.Decode(&apartment)
		apartments = append(apartments, apartment)
	}
	json.NewEncoder(res).Encode(apartments)
}

func GetApartment(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")

	var apartment Apartment
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("StockSpace").Collection("apartments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, Apartment{ID: id}).Decode(&apartment)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(res).Encode(apartment)
}

func DeleteApartment(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("StockSpace").Collection("apartments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collection.DeleteOne(ctx, Apartment{ID: id})
	result := ReqStatus{}
	if err == nil {
		result = ReqStatus{
			Status: "Success!",
			Code:   200,
		}
	} else {
		result = ReqStatus{
			Status: "Fail!",
			Code:   500,
		}
	}
	json.NewEncoder(res).Encode(result)
}

func UpdateApartment(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")

	var apartment Apartment
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	json.NewDecoder(req.Body).Decode(&apartment)
	collection := client.Database("StockSpace").Collection("apartments")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	update := bson.D{primitive.E{
		Key: "$set", Value: bson.D{
			primitive.E{
				Key: "name", Value: apartment.Name},
			primitive.E{
				Key: "unit", Value: apartment.Unit},
			primitive.E{
				Key: "city", Value: apartment.City}}}}
	_, err := collection.UpdateOne(ctx, Apartment{ID: id}, update)

	result := ReqStatus{}
	if err == nil {
		result = ReqStatus{
			Status: "Success!",
			Code:   200,
		}
	} else {
		result = ReqStatus{
			Status: "Fail!",
			Code:   500,
		}
	}
	json.NewEncoder(res).Encode(result)
}

var client *mongo.Client

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println("Server Start!")

	router := mux.NewRouter()
	router.HandleFunc("/apartment", CreateApartment).Methods("POST")
	router.HandleFunc("/apartment", GetAllApartment).Methods("GET")
	router.HandleFunc("/apartment/{id}", GetApartment).Methods("GET")
	router.HandleFunc("/apartment/{id}", DeleteApartment).Methods("DELETE")
	router.HandleFunc("/apartment/{id}", UpdateApartment).Methods("PUT")
	http.ListenAndServe(":3005", router)
}
