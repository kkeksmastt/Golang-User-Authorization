package database

import (
	sl "UserAuth/serverLog"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RefreshToken struct {
	Guid    string    `bson:"_id"`
	Refresh string    `bson:"refresh"`
	Time    time.Time `bson:"time"`
}

var collection *mongo.Collection

func InitDB() {

	MongoDb := os.Getenv("MONGODB_URL")
	if MongoDb == "" {
		MongoDb = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoDb))
	if err != nil {
		sl.ErrorFatal(err)
	}
	collection = client.Database("auth").Collection("refresh-tokens")

	fmt.Println("Connected to MongoDB!")

}

func InsertRefreshToken(refresh, guid string) error {
	now := time.Now()
	etime := now.Add(60 * 24 * time.Hour)

	token := RefreshToken{Guid: guid,
		Refresh: refresh,
		Time:    etime}

	_, err := collection.InsertOne(context.TODO(), token)
	if err != nil {
		return err
	}

	return nil
}

func ReadRefreshToken(guid string) (*RefreshToken, error) {

	filter := bson.D{{Key: "_id", Value: guid}}
	var result *RefreshToken

	var err error
	if err = collection.FindOne(context.TODO(), filter).Decode(&result); err == nil {
		return result, err
	}

	return nil, err
}

func UpdateRefreshToken(refresh, guid string) error {

	filter := bson.D{{Key: "_id", Value: guid}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "refresh", Value: refresh},
			{Key: "time", Value: time.Now().Add(60 * 24 * time.Hour)},
		}},
	}

	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}
	return nil
}
