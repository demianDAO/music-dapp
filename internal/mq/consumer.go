package mq

import (
	"github.com/sirupsen/logrus"
)

func Consume(r *RabbitMQ, queueName string, handler func([]byte) error) {
	logInstance := logrus.WithFields(logrus.Fields{
		"module": "mq",
		"func":   "Consume",
	})

	msgs, err := r.channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		logInstance.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for msg := range msgs {
			err := handler(msg.Body)
			if err != nil {
				logInstance.Errorf("Failed to handle message: %s", err)
			}
		}
	}()
}
