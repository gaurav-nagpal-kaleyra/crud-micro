package main

import (
	"encoding/json"
	"firstExercise/config"
	"math"

	Consumer "firstExercise/consumer"

	health "firstExercise/web/health"
	user "firstExercise/web/userHandlers"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/divan/num2words"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
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
	// if err := config.MySqlConnection(); err != nil {
	// 	zap.L().Fatal("There is some error connecting to mysqlDb", zap.Error(err))
	// }

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

	type NetsuiteInvoiceData struct {
		BillingID string `json:"billing_id" bson:"billing_id"`
		CompanyID string `json:"company_id" bson:"company_id"`
		// OrderID sometimes is provided as string, sometimes as int
		OrderID     interface{} `json:"order_id" bson:"order_id"`
		BillingType int         `json:"billing_type" bson:"billing_type"`
		PaymentMode int         `json:"payment_mode" bson:"payment_mode"`
		PaymentID   string      `json:"payment_id" bson:"payment_id"`
		// Recharge actual value without tax
		BaseValue interface{} `json:"base_value" bson:"base_value"`
		// Recharge value with tax
		GrandTotal interface{} `json:"grand_total" bson:"grand_total"`
		CountryID  int         `json:"country_id" bson:"country_id"`
	}

	type InvoiceData struct {
		OrderID   interface{} `json:"order_id" bson:"order_id"`
		CountryID int         `json:"country_id" bson:"country_id"`
	}

	payload := &InvoiceData{
		OrderID:   4,
		CountryID: 298,
	}

	rmqBody, _ := json.Marshal(payload)
	if err := PublishToRMQ("billing-low-load-queue", rmqBody, "invoice"); err != nil {
		zap.L().Error("error publishing", zap.Error(err))
	}

	fmt.Printf("Server started on port %v\n", os.Getenv("APP_PORT"))

	zap.L().Info(fmt.Sprintf("Listening and Serving on : %s", os.Getenv("APP_PORT")))

	_ = http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("APP_PORT")), router)

}

func PublishToRMQ(qName string, rmqBody []byte, publishType string) error {
	zap.L().Debug(fmt.Sprintf("Publish to %s queue", qName))

	// publish to rmq
	err := config.RMQChan.Publish(
		"",    // exchange
		qName, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        rmqBody,
			Type:        publishType,
		})

	if err != nil {
		return err
	}

	// num to words packagate test
	// getAmountInWords(1000.50)
	convertHTMLtoPDF()
	return nil
}

func getAmountInWords(baseValue float64) {
	intpart, div := math.Modf(baseValue)

	var isNegative bool
	if baseValue < 0 {
		isNegative = true
		intpart = -intpart
	}

	unit := num2words.Convert(int(intpart))
	zap.L().Info("div value", zap.Any("", int(div)))
	divStr := fmt.Sprintf("%.7g", baseValue-float64(intpart))[2:]
	subUnit := num2words.Convert(int(div))
	zap.L().Info("subunit value", zap.Any("", subUnit))
	if isNegative {
		fmt.Printf("(-) %s Rupee and %s paise", unit, divStr)
		return
	}
	fmt.Printf("%s Rupee and %s paise", unit, divStr)
}

func convertHTMLtoPDF() {

}
