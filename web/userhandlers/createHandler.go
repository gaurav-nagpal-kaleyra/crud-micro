package userhandlers

import (
	"encoding/json"
	"firstExercise/config"
	userModel "firstExercise/model/user"
	"firstExercise/redis"
	"firstExercise/repository"

	"net/http"

	"go.uber.org/zap"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user userModel.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		zap.L().Error("Unable to decode request body ", zap.Error(err))
	}

	// sql
	// userRepo := repository.UserRepository{
	// 	Db: config.DB,
	// }

	// userRepo.AddUserInDB(user)

	// store into redis
	if err := redis.AddIntoDBRedis(user); err != nil {
		zap.L().Error("Error inserting into RedisDB", zap.Error(err))
	}

	// mongodb
	mongoRepo := repository.MongoRepository{
		Client: config.Client,
	}
	err = mongoRepo.AddUserInDB(user)

	var resp userModel.Response
	if err != nil {
		resp = userModel.Response{
			StatusCode: 500,
			Error:      "Error creating the user",
			Message:    "Internal Server Error",
			Data:       &user,
		}
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		resp = userModel.Response{
			StatusCode: 200,
			Error:      "",
			Message:    "User Created",
			Data:       &user,
		}
		w.WriteHeader(http.StatusCreated)
	}
	zap.L().Debug("Called the AddUser service")

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		zap.L().Error("Unable to encode responses body ", zap.Error(err))
	}

}
