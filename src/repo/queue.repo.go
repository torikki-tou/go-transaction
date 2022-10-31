package repo

import (
	"context"
	"encoding/json"
	"github.com/torikki-tou/go-transaction/config"
	"github.com/torikki-tou/go-transaction/dto"
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
)

type QueueRepository interface {
	ProduceNotification(clientID string, delta int, changedBalance int) error
}

type queueRepository struct {
	connection *ampq.Connection
}

func NewProducer(connection *ampq.Connection) QueueRepository {
	return &queueRepository{connection: connection}
}

func (p *queueRepository) ProduceNotification(clientID string, delta int, changedBalance int) error {
	ch, err := p.connection.Channel()
	if err != nil {
		return err
	}
	defer func(ch *ampq.Channel) { _ = ch.Close() }(ch)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(dto.Notification{
		ClientID:       clientID,
		Delta:          delta,
		ChangedBalance: changedBalance,
	})
	if err != nil {
		return err
	}
	err = ch.PublishWithContext(ctx,
		"",
		config.QueueName,
		false,
		false,
		ampq.Publishing{
			DeliveryMode: ampq.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		return err
	}
	return nil
}
