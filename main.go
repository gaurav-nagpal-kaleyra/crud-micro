package main

import (
	"crud-micro/config"
	"crud-micro/middleware"

	Consumer "crud-micro/consumer"

	health "crud-micro/web/health"
	user "crud-micro/web/userhandlers"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {

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
		zap.L().Fatal("unable to connect redis ", zap.Error(err))
	}

	// rabbit mq connection
	err = config.InitRabbitMQ()
	if err != nil {
		zap.L().Fatal("Unable to connect RabbitMQ")
	}

	api := flag.Bool("api", false, "For running server/publisher")
	usersConsumer := flag.Bool("users-consumer", false, "For running consumer")
	deleteConsumer := flag.Bool("delete-consumer", false, "For running consumer")
	flag.Parse()

	if *api {
		runHTTPServer()
	}

	if *usersConsumer {
		Consumer.ConsumeMessagesUsersQueue()
		return
	}
	if *deleteConsumer {
		Consumer.ConsumeMessagesDeleteQueue()
		return
	}

}

func runHTTPServer() {
	router := mux.NewRouter()
	r := router.PathPrefix("/").Subrouter()
	router.HandleFunc("/v1/health", health.HealthHandler).Methods("GET")
	r.Use(middleware.JwtMiddleware)

	// CRUD endpoints start
	router.HandleFunc("/v1/user/create", user.CreateHandler).Methods("POST")
	r.HandleFunc("/v1/user/read", user.ReadHandler).Methods("GET")
	r.HandleFunc("/v1/user/update", user.UpdateHandler).Methods("PUT")
	r.HandleFunc("/v1/user/delete", user.DeleteHandler).Methods("DELETE")
	// CRUD endpoints end
	fmt.Printf("Server started on port %v\n", os.Getenv("APP_PORT"))

	zap.L().Info(fmt.Sprintf("Listening and Serving on : %s", os.Getenv("APP_PORT")))

	_ = http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("APP_PORT")), router)
}
