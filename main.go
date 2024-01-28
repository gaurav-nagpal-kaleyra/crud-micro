package main

import (
	"crud-micro/config"

	Consumer "crud-micro/consumer"

	health "crud-micro/web/health"
	user "crud-micro/web/userHandlers"
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
	consumer := flag.Bool("consumer", false, "For running consumer")
	flag.Parse()

	if *api {
		runHTTPServer()
	}

	if *consumer {
		Consumer.ConsumeMessages()
	}

}

func runHTTPServer() {
	router := mux.NewRouter()

	router.HandleFunc("/v1/health", health.HealthHandler).Methods("GET")
	router.HandleFunc("/v1/user/create", user.CreateHandler).Methods("POST")
	router.HandleFunc("/v1/user/read/", user.ReadHandler).Methods("GET")

	fmt.Printf("Server started on port %v\n", os.Getenv("APP_PORT"))

	zap.L().Info(fmt.Sprintf("Listening and Serving on : %s", os.Getenv("APP_PORT")))

	_ = http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("APP_PORT")), router)

}
