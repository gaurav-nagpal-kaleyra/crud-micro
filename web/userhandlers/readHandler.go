package userhandlers

import (
	"encoding/json"
	"firstExercise/config"
	usermodel "firstExercise/model/user"
	"firstExercise/repository"

	"net/http"

	"go.uber.org/zap"
)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()

	userId := queryParams["userId"][0]

	zap.L().Debug("Calling the FindUser Service")

	// for sql
	// userRepo := repository.UserRepository{
	// 	Db: config.DB,
	// }

	// userFound := userRepo.FindUserFromDB(userId)

	// for mongodb
	mongoRepo := repository.MongoRepository{
		Client:     config.Client,
		Collection: config.Collection,
	}

	userFound := mongoRepo.FindUserFromDB(userId)

	var resp usermodel.Response
	resp = usermodel.Response{
		StatusCode: 200,
		Error:      "",
		Message:    "User Found",
		Data:       &userFound,
	}

	if userFound.UserName == "" {

		resp = usermodel.Response{
			StatusCode: 200,
			Error:      "",
			Message:    "User Not Found",
			Data:       nil,
		}

	}
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		zap.L().Error("Unabel to encode response body", zap.Error(err))
	}

}
