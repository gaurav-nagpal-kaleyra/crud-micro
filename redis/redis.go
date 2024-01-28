package redis

import (
	"crud-micro/config"
	userModel "crud-micro/model/user"
	"encoding/json"
	"fmt"
	"strconv"
)

// can be also used to update user info
func AddIntoDBRedis(newUser *userModel.User) error {
	// storing in redis string
	userId := strconv.Itoa(newUser.UserId)
	json, err := json.Marshal(newUser)
	if err != nil {
		return err
	}
	return config.RedisConn.Set(fmt.Sprintf("usersInfo:%s", userId), json, 0).Err()
}

func ReadFromDBRedis(userId string) (*userModel.User, error) {
	var u userModel.User

	val, err := config.RedisConn.Get(fmt.Sprintf("usersInfo:%s", userId)).Result()

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(val), &u)

	return &u, err
}

func DeleteFromDBRedis(userId string) error {
	key := fmt.Sprintf("usersInfo:%s", userId)
	return config.RedisConn.Del(key).Err()
}
