package config

import ampq "github.com/rabbitmq/amqp091-go"

var QueueName = "notifications"

func InitQueue(con *ampq.Connection) {
	ch, err := con.Channel()
	if err != nil {
		panic(err)
	}
	defer func(ch *ampq.Channel) { _ = ch.Close() }(ch)

	_, err = ch.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
}

func SetupRabbitMQConnection() *ampq.Connection {
	con, err := ampq.Dial("amqp://guest:guest@rabbit:5672/")
	if err != nil {
		panic(err)
	}
	return con
}

func CloseRabbitMQConnection(con *ampq.Connection) {
	_ = con.Close()
}
