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

/*
For the first exercise you're expected to
create a Golang project (please use Golang v1.18)
create a mux server
add a logger that prints significant information (we use zap logger)
the data will be hardcoded for now, no db connection is required
add unit tests for the endpoints

Here's a list of few improvements you can make :slightly_smiling_face:

-have the project structured in folders
-the models should stay in a separate package (you can either call it dto or models as you prefer)
-move the User model under a separate package
-consider having a model for your responses that contains a success/error message along with the returned data

-the business logic should stay in a separate package

-this means that the endpoint handlers will call some methods from the service or user package to get the data
-the business logic (append to the users list), will happen there

consider having a config package
it can be useful to have specific values configured, (for instance the port number where the service will run)
it can be useful for initialization of all the required items at startup
for instance the zap logger, when you'll add it :slightly_smiling_face:

for your endpoints

-consider always having a health endpoint, that will return the status of your server (it can be a simple json object containing {"status": "OK"})

consider always versioning your endpoints, so that if you need to add a new functionality for which the response returned is different, you'll create a new endpoint, similar to the previous one, with an increased version number. This can be achieved by simply adding a /v1/ prefix to your endpoints signature
you can also consider having the handlers in a separate package (for instance web or api)
*/

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

	router.HandleFunc("/v1/health", health.HealthHandler).Methods("GET")
	router.HandleFunc("/v1/user/create", user.CreateHandler).Methods("POST")
	router.HandleFunc("/v1/user/read/", user.ReadHandler).Methods("GET")

	fmt.Printf("Server started on port %v\n", os.Getenv("APP_PORT"))

	zap.L().Info(fmt.Sprintf("Listening and Serving on : %s", os.Getenv("APP_PORT")))

	err := http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("APP_PORT")), router)

	if err != nil {
		zap.L().Error("Listening and Serving Error", zap.Error(err))
		return
	}

}
