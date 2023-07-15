package userservice

import (
	usersDb "firstExercise/Db"
	userModel "firstExercise/model/user"
	"strconv"
)

func FindUser(userId string) userModel.User {
	var user userModel.User

	userIdInt, _ := strconv.Atoi(userId)

	for _, u := range usersDb.UserList {
		if u.ID == userIdInt {
			return u
		}
	}

	//returning empty struct if user not found
	return user
}
