package userhandlers

import (
	"crud-micro/constant"
	userModel "crud-micro/model/user"
	"crud-micro/rabbitmq"
	"crud-micro/redis"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	// first see if it exists in the redis, otherwise raise error
	userID := r.URL.Query().Get("userId")

	var user *userModel.User

	user, err := redis.ReadFromDBRedis(userID)
	if err != nil {
		zap.L().Error("error reading from redis", zap.Error(err))
		resp := userModel.Response{
			StatusCode: 500,
			Error:      "Error getting the user info from redis",
			Message:    "Internal Server Error",
			Data:       nil,
		}
		w.WriteHeader(http.StatusInternalServerError)

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			zap.L().Error("Unable to encode response body ", zap.Error(err))
		}
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		zap.L().Error("Unable to decode request body ", zap.Error(err))
	}

	// update in redis - now it means user exists in redisDB
	if err := redis.AddIntoDBRedis(user); err != nil {
		zap.L().Error("error updating user in redis", zap.Error(err))
		resp := userModel.Response{
			StatusCode: 500,
			Error:      "Error deleting the user info from redis",
			Message:    "Internal Server Error",
			Data:       nil,
		}
		w.WriteHeader(http.StatusInternalServerError)

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			zap.L().Error("Unable to encode responses body ", zap.Error(err))
		}
		return
	}

	// push it into rmq to update in mysql and mongodb
	// push userID in usersqueue
	rmqBody, err := json.Marshal(user)
	if err != nil {
		zap.L().Error("Publish To Queue - Error in Marshalling")
		return
	}

	err = rabbitmq.PublishToQueue(os.Getenv(constant.UsersQueue), rmqBody, "")

	if err != nil {
		zap.L().Error(fmt.Sprintf("Error in publishing the message to %s queue", os.Getenv(constant.UsersQueue)))
	}

	resp := userModel.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "Success",
		Data:       nil,
	}
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		zap.L().Error("Unable to encode responses body ", zap.Error(err))
	}
}
