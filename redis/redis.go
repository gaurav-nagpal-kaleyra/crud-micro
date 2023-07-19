package redis

import (
	"encoding/json"
	"firstExercise/config"
	userModel "firstExercise/model/user"
	"strconv"
)

func AddIntoDBRedis(newUser userModel.User) error {

	// storing in redis hash
	userId := strconv.Itoa(newUser.UserId)
	json, err := json.Marshal(newUser)
	if err != nil {
		return err
	}
	return config.RedisConn.HSet("usersInfo", userId, json).Err()

}

func ReadFromDBRedis(userId string) (userModel.User, error) {
	var u userModel.User

	val, err := config.RedisConn.HGet("usersInfo", userId).Result()
	if err != nil {
		return u, err
	}

	err = json.Unmarshal([]byte(val), &u)

	if err != nil {
		return u, err
	}

	return u, nil
}
