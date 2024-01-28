package consumer

import (
	"crud-micro/config"
	"crud-micro/constant"
	userModel "crud-micro/model/user"
	"crud-micro/repository"
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

// consumer to consume the messages from the users_queue
func ConsumeMessagesUsersQueue() {
	msgs, err := config.RMQChan.Consume(
		os.Getenv(constant.UsersQueue), // queue name
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		zap.L().Error("Error in consuming the message", zap.Error(err))
	}
	userRepo := repository.UserRepository{
		Db: config.DB,
	}
	mongoRepo := repository.MongoRepository{
		Client: config.Client,
	}
	for msg := range msgs {
		// add into mysql and mongodb
		var user *userModel.User
		if err := json.Unmarshal(msg.Body, &user); err != nil {
			zap.L().Error("error unmarshaling user msg", zap.Error(err))
			continue
		}
		zap.L().Info("message received", zap.Any("", user))

		// create user in mysql and mongodb
		if err := userRepo.AddUserInDB(user); err != nil {
			zap.L().Error("error inserting user in mysql", zap.Error(err))
		}
		if err := mongoRepo.AddUserInDB(user); err != nil {
			zap.L().Error("error inserting user in mysql", zap.Error(err))
		}
	}
}

// consumer to consume messages from the delete queue
func ConsumeMessagesDeleteQueue() {
	msgs, err := config.RMQChan.Consume(
		os.Getenv(constant.DeleteUsersQueue), // queue name
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		zap.L().Error("Error in consuming the message", zap.Error(err))
	}

	userRepo := repository.UserRepository{
		Db: config.DB,
	}
	mongoRepo := repository.MongoRepository{
		Client: config.Client,
	}
	for msg := range msgs {
		// add into mysql and mongodb

		var userId string
		if err := json.Unmarshal(msg.Body, &userId); err != nil {
			zap.L().Error("error unmarshaling user msg", zap.Error(err))
			continue
		}
		zap.L().Info("message received", zap.Any("", userId))

		// delete the users with the userID from mongodb and mysql
		if err := userRepo.DeleteUserFromDB(userId); err != nil {
			zap.L().Error("error deleting from mysql", zap.Error(err))
		}

		if err := mongoRepo.DeleteUserFromDB(userId); err != nil {
			zap.L().Error("error deleting from mysql", zap.Error(err))
		}
	}
}
