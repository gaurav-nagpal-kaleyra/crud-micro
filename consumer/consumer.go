package consumer

import (
	"crud-micro/config"
	"fmt"
	"os"

	"go.uber.org/zap"
)

// consumer to consume the messages from the users_queue

func ConsumeMessages() {
	msgs, err := config.RMQChan.Consume(
		os.Getenv("USERS_QUEUE"), // queue name
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		zap.L().Error("Error in consuming the message", zap.Error(err))
	}

	for msg := range msgs {
		zap.L().Info(fmt.Sprintf("Message Received :%s \n", msg.Body))
	}

}
