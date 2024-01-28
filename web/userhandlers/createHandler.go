package userhandlers

import (
	"crud-micro/config"
	"crud-micro/constant"
	userModel "crud-micro/model/user"
	"crud-micro/rabbitmq"
	"crud-micro/redis"
	"crud-micro/repository"
	"crud-micro/utils"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"net/http"

	"go.uber.org/zap"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var user userModel.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		zap.L().Error("Unable to decode request body ", zap.Error(err))
	}
	// when the user is created, we generate a JWT token
	token, err := utils.CreateJwtToken(strconv.Itoa(user.UserId))
	if err != nil {
		zap.L().Error("error creating jwt token", zap.Error(err))
		resp := userModel.Response{
			StatusCode: 500,
			Error:      "Error creating the token",
			Message:    "Internal Server Error",
			Data:       &user,
		}
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			zap.L().Error("Unable to encode responses body ", zap.Error(err))
		}
		return
	}

	w.Header().Set("token", token)
	// sql
	// userRepo := repository.UserRepository{
	// 	Db: config.DB,
	// }

	// userRepo.AddUserInDB(user)

	// store into redis
	if err := redis.AddIntoDBRedis(&user); err != nil {
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

		// whenever the user is created, message is published to the users_queue
		rmqBody, err := json.Marshal("User Added")
		if err != nil {
			zap.L().Error("Publish To Queue - Error in Marshalling")
			return
		}

		err = rabbitmq.PublishToQueue(os.Getenv(constant.UsersQueue),
			rmqBody, "")

		if err != nil {
			zap.L().Error(fmt.Sprintf("Error in publishing the message to %s queue", os.Getenv("USERS_QUEUE")))
		}
	}
	zap.L().Debug("Called the AddUser service")

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		zap.L().Error("Unable to encode responses body ", zap.Error(err))
	}

}
