package redis

import (
	"encoding/json"
	"firstExercise/config"
	userModel "firstExercise/model/user"
	"fmt"
	"strconv"
)

func AddIntoDBRedis(newUser *userModel.User) error {

	// storing in redis string
	userId := strconv.Itoa(newUser.UserId)
	json, err := json.Marshal(newUser)
	if err != nil {
		return err
	}
	return config.RedisConn.Set(fmt.Sprintf("usersInfo:%s", userId), json, 0).Err()

}

func ReadFromDBRedis(userId string) (userModel.User, error) {
	var u userModel.User

	val, err := config.RedisConn.Get(fmt.Sprintf("usersInfo:%s", userId)).Result()

	err = json.Unmarshal([]byte(val), &u)

	return u, err
}
