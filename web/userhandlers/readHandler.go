package userhandlers

import (
	"encoding/json"
	"firstExercise/config"
	usermodel "firstExercise/model/user"
	"firstExercise/redis"
	"firstExercise/repository"

	"net/http"

	"go.uber.org/zap"
)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()

	userId := queryParams["userId"][0]

	var userFound usermodel.User

	// first check in the redis cache
	res, err := redis.ReadFromDBRedis(userId)
	if err == nil {
		userFound = res
		zap.L().Info("Get from redisDB", zap.Any("User", userFound))
	} else {
		zap.L().Debug("Calling the FindUser Service")

		// for sql
		// userRepo := repository.UserRepository{
		// 	Db: config.DB,
		// }

		// userFound := userRepo.FindUserFromDB(userId)

		// for mongodb

		mongoRepo := repository.MongoRepository{
			Client: config.Client,
		}

		userFound = mongoRepo.FindUserFromDB(userId)
	}

	var resp usermodel.Response

	// if the user is not found
	if userFound.UserName == "" {
		resp = usermodel.Response{
			StatusCode: 200,
			Error:      "",
			Message:    "User Not Found",
			Data:       nil,
		}
	} else {
		resp = usermodel.Response{
			StatusCode: 200,
			Error:      "",
			Message:    "User Found",
			Data:       &userFound,
		}
	}
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(resp)

	if err != nil {
		zap.L().Error("Unabel to encode response body", zap.Error(err))
	}

}
