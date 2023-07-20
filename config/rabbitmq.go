package config

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

var RMQConn *amqp.Connection
var RMQChan *amqp.Channel

func InitRabbitMQ() error {

	if RMQConn == nil || RMQConn.IsClosed() {
		RMQConn, err = rabbitMQConnection()
		if err != nil {
			return err
		}
	}
	// create rabbitmq channel
	RMQChan, err = RMQConn.Channel()
	if err != nil {
		return err
	}

	zap.L().Info("new RMQ connection and channel successfully created.")

	// declare required queues
	return declareQueues()
}

func rabbitMQConnection() (*amqp.Connection, error) {
	Host := os.Getenv("RABBITMQ_HOST")
	Port := os.Getenv("RABBITMQ_PORT")
	Username := os.Getenv("RABBITMQ_USER")
	Password := os.Getenv("RABBITMQ_PASS")

	connectionString := fmt.Sprintf("amqp://%s:%s@%s:%s", Username, Password, Host, Port)

	conn, err := amqp.Dial(connectionString)

	if err != nil {
		return nil, err
	}

	zap.L().Info("RabbitMQ Successfully Connected.")

	return conn, nil
}

func declareQueues() error {
	// declare the "users_queue"
	_, err := RMQChan.QueueDeclare(
		"users_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		zap.L().Error("Error declaring the users_queue", zap.Error(err))
		return err
	}
	return nil
}
