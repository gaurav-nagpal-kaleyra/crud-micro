package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

var DB *sql.DB
var err error

func MySqlConnection() error {
	const (
		username = "root"
		password = "tiger2001"
		hostname = "127.0.0.1:3306"
		dbname   = "usersDB"
	)
	// "mysql", "root:tiger2001@tcp(127.0.0.1:3306)/usersDB"
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname))

	if err != nil {
		zap.L().Error("Error while establishing connection to database", zap.Error(err))
		return err
	}

	if err := DB.Ping(); err != nil {
		zap.L().Error("Error while verifying database conection", zap.Error(err))
		return err
	}

	zap.L().Info("Mysql connection established")

	// creating table for storing user data
	if err := createUserInfoTable(DB); err != nil {
		zap.L().Error("Unable to create the table ", zap.Error(err))
		return err
	}

	return nil
}

func createUserInfoTable(DB *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS userInfo(id int primary key auto_increment, userName text, userAge int, userLocation text)`

	res, err := DB.Exec(query)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("Rows affected", rows)
	return nil
}
