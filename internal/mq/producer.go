package mq

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func Publish(r *RabbitMQ, routingKey string, message []byte) error {
	var logInstance = log.WithFields(log.Fields{
		"module": "mq",
		"func":   "Publish",
	})
	err := r.channel.Publish(
		ExchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	if err != nil {
		logInstance.Printf("Failed to publish a message: %s", err)
		return err
	}
	return nil
}
