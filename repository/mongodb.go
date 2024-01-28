package repository

import (
	"context"
	"firstExercise/model"
	userModel "firstExercise/model/user"
	"fmt"
	"strconv"
	"time"

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
func (m *MongoRepository) AddUserInDB(newUser userModel.User) error {
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

func (m *MongoRepository) FindUserFromDB(userId string) userModel.User {
	// find from mongodb
	m.Collection = check(m)
	var u userModel.User
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		zap.L().Error("Error converting string to int", zap.Error(err))
	}

	filter := bson.D{{Key: "userId", Value: userIdInt}}
	err = m.Collection.FindOne(context.Background(), filter).Decode(&u)
	zap.L().Debug("FindOne function called")
	if err != nil {
		zap.L().Error("Error finding the user", zap.Error(err))
	}

	return u
}

func (m *MongoRepository) GetBTLDocumentCountBasedOnDate(startDate, endDate time.Time, companyID string) (int, error) {
	m.Collection = m.Client.Database("usersDB").Collection("balance_transaction_log_testcompany")
	filter := bson.D{
		{
			Key: "date",
			Value: bson.D{
				{Key: "$gte", Value: startDate},
				{Key: "$lte", Value: endDate},
			},
		},
	}
	dbCount, err := m.Collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return int(dbCount), nil
}

func (m *MongoRepository) FetchBtlDataForDumping(startDate, endDate time.Time, companyID string) ([]*model.BalanceTransaction, error) {
	m.Collection = m.Client.Database("usersDB").Collection("balance_transaction_log_testcompany")
	// filter := bson.D{
	// 	{
	// 		Key: "date",
	// 		Value: bson.D{
	// 			{Key: "$gte", Value: startDate},
	// 			{Key: "$lte", Value: endDate},
	// 		},
	// 	},
	// }
	// pipeline := mongo.Pipeline{
	// 	// add all the filters
	// 	{{Key: "$match", Value: filter}},

	// 	// sort on the basis of date
	// 	{{Key: "$sort", Value: bson.D{
	// 		{Key: "date", Value: 1},
	// 	}}},
	// }

	btlData := []*model.BalanceTransaction{}

	// cur, err := m.Collection.Aggregate(context.Background(), pipeline)
	// if err != nil {
	// 	return nil, err
	// }

	// if err := cur.All(context.Background(), &btlData); err != nil {
	// 	return nil, err
	// }
	return btlData, nil
}
