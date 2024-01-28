package userhandlers

import (
	"crud-micro/config"
	userModel "crud-micro/model/user"
	"encoding/json"

	"crud-micro/redis"
	"crud-micro/repository"
	"net/http"

	"go.uber.org/zap"
)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	var userFound *userModel.User

	// first check in the redis cache
	res, err := redis.ReadFromDBRedis(userId)
	if err == nil {
		userFound = res
		zap.L().Info("Get from redisDB", zap.Any("User: ", userFound))
	} else {
		zap.L().Debug("Calling the FindUser Service")

		// for sql
		userRepo := repository.UserRepository{
			Db: config.DB,
		}

		userFound = userRepo.FindUserFromDB(userId)

		// for mongodb - if not found in mysql
		if userFound == nil {
			mongoRepo := repository.MongoRepository{
				Client: config.Client,
			}

			userFound = mongoRepo.FindUserFromDB(userId)
		}
	}

	var resp userModel.Response

	// if the user is not found
	if userFound.UserName == "" {
		resp = userModel.Response{
			StatusCode: 200,
			Error:      "",
			Message:    "User Not Found",
			Data:       nil,
		}
	} else {
		resp = userModel.Response{
			StatusCode: 200,
			Error:      "",
			Message:    "User Found",
			Data:       userFound,
		}
	}
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		zap.L().Error("Unabel to encode response body", zap.Error(err))
	}
}
