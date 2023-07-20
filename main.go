package main

import (
	"firstExercise/config"
	health "firstExercise/web/health"
	user "firstExercise/web/userHandlers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	router := mux.NewRouter()

	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading the .env file")
	}

	// initializing the logger
	if err := config.InitializeLogger(); err != nil {
		log.Fatal("Unable to start logger ", err)
	}
	zap.L().Info(".env file load completed")

	// mysql connection
	if err := config.MySqlConnection(); err != nil {
		zap.L().Fatal("There is some error connecting to mysqlDb", zap.Error(err))
	}

	// mongodb connection
	if err := config.MongoDBConnection(); err != nil {
		zap.L().Fatal("There is some error connecting to mongodb", zap.Error(err))
	}

	// redis connection
	err := config.RedisConnection()
	if err != nil {
		zap.L().Fatal("unable to connect redis ",
			zap.Error(err))
	}

	// rabbit mq connection

	router.HandleFunc("/v1/health", health.HealthHandler).Methods("GET")
	router.HandleFunc("/v1/user/create", user.CreateHandler).Methods("POST")
	router.HandleFunc("/v1/user/read/", user.ReadHandler).Methods("GET")

	fmt.Printf("Server started on port %v\n", os.Getenv("APP_PORT"))

	zap.L().Info(fmt.Sprintf("Listening and Serving on : %s", os.Getenv("APP_PORT")))

	err = http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("APP_PORT")), router)

	if err != nil {
		zap.L().Error("Listening and Serving Error", zap.Error(err))
		return
	}

}
