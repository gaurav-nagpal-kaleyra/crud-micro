package userhandlers

import (
	"encoding/json"
	userModel "firstExercise/model/user"
	userService "firstExercise/service/userService"
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

	zap.L().Debug("Called the AddUser service")
	userService.AddUser(user)
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
