package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

var Client *mongo.Client

func MongoDBConnection() error {

	Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))))
	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}
	zap.L().Info("Mongodb connection established")

	return nil
}
