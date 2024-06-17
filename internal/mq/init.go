package mq

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"web3-music-platform/internal/app/song/handlers"
)

var logIns = logrus.WithFields(logrus.Fields{
	"pkg":  "mq",
	"func": "Init",
})

var RabbitMQInstance *RabbitMQ

const (
	ExchangeName      = "song"
	SongUploadQueue   = "song_upload_queue"
	CreateSongQueue   = "create_song_queue"
	PurchaseSongQueue = "purchase_song_queue"
)

var Queues = map[string]string{
	SongUploadQueue:   "song.upload",
	CreateSongQueue:   "song.create",
	PurchaseSongQueue: "song.purchase",
}

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func Init(amqpURL string) error {

	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		logIns.WithError(err).Error("Failed to connect to RabbitMQ")
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		logIns.WithError(err).Error("Failed to open a channel")
		return err
	}

	err = ch.ExchangeDeclare(
		ExchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		logIns.WithError(err).Error("Failed to declare an exchange")
		return err
	}

	for queue, routingKey := range Queues {
		q, err := ch.QueueDeclare(
			queue,
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		if err != nil {
			logIns.WithError(err).Errorf("Failed to declare queue %s", queue)
			return err
		}

		err = ch.QueueBind(q.Name, routingKey, ExchangeName, false, nil)
		if err != nil {
			logIns.WithError(err).Errorf("Failed to bind queue %s", queue)
			return err
		}
	}

	RabbitMQInstance = &RabbitMQ{conn: conn, channel: ch}

	go Consume(RabbitMQInstance, SongUploadQueue, handlers.SongUpload)
	go Consume(RabbitMQInstance, CreateSongQueue, handlers.CreateMusicNFT)
	go Consume(RabbitMQInstance, PurchaseSongQueue, handlers.PurchaseSong)

	return nil
}

func (r *RabbitMQ) Close() {
	err := r.channel.Close()
	if err != nil {
		logIns.WithError(err).Error("Failed to close channel")
	}
	err = r.conn.Close()
	if err != nil {
		logIns.WithError(err).Error("Failed to close connection")
	}
}
