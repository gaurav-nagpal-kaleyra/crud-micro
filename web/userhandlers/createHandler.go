package userhandlers

import (
	"encoding/json"
	"firstExercise/config"
	userModel "firstExercise/model/user"
	"firstExercise/repository"

	"net/http"

	"go.uber.org/zap"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var user userModel.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		zap.L().Error("Unable to decode request body ", zap.Error(err))
	}

	// from the mysql
	// userRepo := repository.UserRepository{
	// 	Db: config.DB,
	// }

	// userRepo.AddUserInDB(user)

	// from the mongodb
	mongoRepo := repository.MongoRepository{
		Client:     config.Client,
		Collection: config.Collection,
	}
	mongoRepo.AddUserInDB(user)
	zap.L().Debug("Called the AddUser service")

	resp := userModel.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "User Created",
		Data:       &user,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		zap.L().Error("Unable to encode responses body ", zap.Error(err))
	}

}
