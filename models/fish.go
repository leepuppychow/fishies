package models

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/leepuppychow/fishies/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AllFish() ([]primitive.M, error) {
	var results []bson.M
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	collection := database.DB.Database(os.Getenv("DB_NAME")).Collection("fish")
	cur, err := collection.Find(ctx, bson.D{})
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
	}
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
		}
		results = append(results, result)
	}
	return results, err
}

func CreateFish(body io.Reader) (interface{}, error) {
	var data interface{}
	err := json.NewDecoder(body).Decode(&data)
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	collection := database.DB.Database(os.Getenv("DB_NAME")).Collection("fish")
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		log.Println(err)
	}
	return res.InsertedID, err
}
