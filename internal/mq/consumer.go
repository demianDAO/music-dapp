package mq

import log "github.com/sirupsen/logrus"

func Consume(r *RabbitMQ, handler func([]byte) error) {
	var logInstance = log.WithFields(log.Fields{
		"module": "mq",
		"func":   "Consume",
	})

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
		logInstance.Fatalf("Failed to register a consumer: %s", err)
	}

	//var forever = make(chan bool)

	go func() {
		for msg := range msgs {
			if err := handler(msg.Body); err != nil {
				logInstance.Printf("Failed to handle message: %s", err)
			}
		}
	}()

	//<-forever
}
