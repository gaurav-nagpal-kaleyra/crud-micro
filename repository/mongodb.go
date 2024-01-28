package repository

import (
	"context"
	userModel "crud-micro/model/user"
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

func check(m *MongoRepository) *mongo.Collection {
	if m.Collection == nil {
		return m.Client.Database("usersDB").Collection("userInfo")
	}
	return m.Collection
}
func (m *MongoRepository) AddUserInDB(newUser *userModel.User) error {
	// storing into mongodb
	m.Collection = check(m)
	user := bson.D{{Key: "userId", Value: newUser.UserId}, {Key: "userName", Value: newUser.UserName}, {Key: "userAge", Value: newUser.UserAge}, {Key: "userLocation", Value: newUser.UserLocation}}

	result, err := m.Collection.InsertOne(context.Background(), user)
	zap.L().Debug("InsertOne function called")
	if err != nil {
		return err
	}
	zap.L().Info(fmt.Sprintf("User with id %v inserted ", result.InsertedID))
	return nil
}

func (m *MongoRepository) FindUserFromDB(userId string) *userModel.User {
	// find from mongodb
	m.Collection = check(m)
	var u userModel.User
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		zap.L().Error("Error converting string to int", zap.Error(err))
		return nil
	}

	filter := bson.D{{Key: "userId", Value: userIdInt}}
	err = m.Collection.FindOne(context.Background(), filter).Decode(&u)
	zap.L().Debug("FindOne function called")
	if err != nil {
		zap.L().Error("Error finding the user", zap.Error(err))
	}

	return &u
}

func (m *MongoRepository) DeleteUserFromDB(userID string) error {
	m.Collection = check(m)
	userIdInt, _ := strconv.Atoi(userID)
	filter := bson.D{{Key: "userId", Value: userIdInt}}
	_, err := m.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
