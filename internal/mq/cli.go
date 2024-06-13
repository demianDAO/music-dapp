package mq

import (
	"github.com/streadway/amqp"
	"log"
)

var RabbitMQInstance *RabbitMQ

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMQ(amqpURL string) error {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"music_upload_queue",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	RabbitMQInstance = &RabbitMQ{conn: conn, channel: ch, queue: q}

	return nil
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) Publish(message []byte) error {
	err := r.channel.Publish(
		"",
		r.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	if err != nil {
		log.Printf("Failed to publish a message: %s", err)
		return err
	}
	return nil
}

func (r *RabbitMQ) Consume(handler func([]byte) error) {
	msgs, err := r.channel.Consume(
		r.queue.Name,
		"",
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for d := range msgs {
			if err := handler(d.Body); err != nil {
				log.Printf("Failed to handle message: %s", err)
			}
		}
	}()
}
