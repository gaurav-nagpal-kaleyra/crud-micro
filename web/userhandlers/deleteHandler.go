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

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// here we only update in redis and push it into queue
	// consumer deletes in mongodb and mysql

	userId := r.URL.Query().Get("userId")
	if userId == "" {
		zap.L().Debug("userId not passed")

		resp := userModel.Response{
			StatusCode: 200,
			Error:      "",
			Message:    "Please pass userID",
			Data:       nil,
		}
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(resp)
		return
	}

	if err := redis.DeleteFromDBRedis(userId); err != nil {
		zap.L().Debug("error while deleting from redis")

		resp := userModel.Response{
			StatusCode: 500,
			Error:      "",
			Message:    "Internal Server error",
			Data:       nil,
		}
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(resp)
	}

	// push userID in usersqueue
	rmqBody, err := json.Marshal(userId)
	if err != nil {
		zap.L().Error("Publish To Queue - Error in Marshalling")
		return
	}

	err = rabbitmq.PublishToQueue(os.Getenv(constant.DeleteUsersQueue), rmqBody, "")
	if err != nil {
		zap.L().Error(fmt.Sprintf("Error in publishing the message to %s queue", os.Getenv("USERS_QUEUE")))
	}

	if err := rabbitmq.PublishToQueue(constant.UsersQueue, rmqBody, ""); err != nil {
		zap.L().Error("error publishing to rmq", zap.Error(err))
	}

	resp := userModel.Response{
		StatusCode: 500,
		Error:      "",
		Message:    "Deleted Successfully",
		Data:       nil,
	}
	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(resp)
}
