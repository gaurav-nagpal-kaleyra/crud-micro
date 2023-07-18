package repository

import (
	"database/sql"
	userModel "firstExercise/model/user"
	"strconv"

	"go.uber.org/zap"
)

type UserRepository struct {
	Db *sql.DB
}

func (r *UserRepository) AddUserInDB(newUser userModel.User) error {

	// storing into mysql
	query := "INSERT INTO userInfo(id, userName, userAge, userLocation) VALUES(?,?,?,?)"
	_, err := r.Db.Exec(query, newUser.UserId, newUser.UserName, newUser.UserAge, newUser.UserLocation)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindUserFromDB(userId string) userModel.User {

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		zap.L().Error("Error converting into int ", zap.Error(err))
	}
	//find from sql
	query := `SELECT * FROM userInfo where id = ? `
	rows, err := r.Db.Query(query, userIdInt)
	if err != nil {
		zap.L().Error("Error selecting from mysql database", zap.Error(err))
	}
	var u userModel.User

	for rows.Next() {
		rows.Scan(&u.UserId, &u.UserName, &u.UserAge, &u.UserLocation)
	}

	return u
}
