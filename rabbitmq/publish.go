package rabbitmq

import (
	"firstExercise/config"
	"fmt"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func PublishToQueue(qName string, rmqBody []byte, publishType string) error {
	zap.L().Debug(fmt.Sprintf("Publish to %s queue", qName))

	// publish to rmq
	err := config.RMQChan.Publish(
		"", // using default exchange
		qName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        rmqBody,
			Type:        publishType,
		})

	if err != nil {
		zap.L().Error("Error while publishing", zap.String("queueName", qName), zap.Error(err))
		return err
	}

	return nil
}
