package userservice

import (
	userDb "firstExercise/Db"
	userModel "firstExercise/model/user"
)

func AddUser(newUser userModel.User) {
	userDb.UserList = append(userDb.UserList, newUser)
}
