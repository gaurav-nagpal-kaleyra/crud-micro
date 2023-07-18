package repository

import (
	"context"
	userModel "firstExercise/model/user"
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type MongoRepository struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func (m *MongoRepository) AddUserInDB(newUser userModel.User) {
	// storing into mongodb
	user := bson.D{{Key: "userId", Value: newUser.UserId}, {Key: "userName", Value: newUser.UserName}, {Key: "userAge", Value: newUser.UserAge}, {Key: "userLocation", Value: newUser.UserLocation}}

	result, err := m.Collection.InsertOne(context.TODO(), user)
	if err != nil {
		zap.L().Error("Error inserting document into mongodb")
	}

	zap.L().Info(fmt.Sprintf("User with id %v inserted ", result.InsertedID))
}

func (m *MongoRepository) FindUserFromDB(userId string) userModel.User {
	// find from mongodb
	var u userModel.User
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		zap.L().Error("Error converting string to int", zap.Error(err))
	}
	fmt.Println(userIdInt)
	filter := bson.D{{Key: "userId", Value: userIdInt}}
	err = m.Collection.FindOne(context.TODO(), filter).Decode(&u)
	fmt.Println(u)
	if err != nil {
		zap.L().Error("Error finding the user", zap.Error(err))
	}

	return u
}
