package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	ampq "github.com/rabbitmq/amqp091-go"
	"github.com/torikki-tou/go-transaction/config"
	"github.com/torikki-tou/go-transaction/handler/v1"
	"github.com/torikki-tou/go-transaction/repo"
	"github.com/torikki-tou/go-transaction/service"
)

var (
	db            *sql.DB               = config.SetupDatabaseConnection()
	rabbitCon     *ampq.Connection      = config.SetupRabbitMQConnection()
	clientRepo    repo.ClientRepository = repo.NewClientRepository(db)
	queueRepo     repo.QueueRepository  = repo.NewProducer(rabbitCon)
	userService   service.ClientService = service.NewClientService(clientRepo, queueRepo)
	clientHandler v1.ClientHandler      = v1.NewClientHandler(userService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	defer config.CloseRabbitMQConnection(rabbitCon)

	config.InitDB(db)
	config.InitQueue(rabbitCon)

	server := gin.Default()
	apiGroup := server.Group("/api")

	transactionGroup := apiGroup.Group("/client")
	{
		transactionGroup.POST("/change_balance", clientHandler.ChangeBalance)
	}

	_ = server.Run()
}
