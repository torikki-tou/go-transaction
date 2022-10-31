package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/torikki-tou/go-transaction/config"
	"github.com/torikki-tou/go-transaction/handler/v1"
	"github.com/torikki-tou/go-transaction/repo"
	"github.com/torikki-tou/go-transaction/service"
)

var (
	db            *sql.DB             = config.SetupDatabaseConnection()
	userRepo      repo.UserRepository = repo.NewUserRepository(db)
	userService   service.UserService = service.NewUserService(userRepo)
	clientHandler v1.ClientHandler    = v1.NewClientHandler(userService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	config.InitDB(db)

	server := gin.Default()
	apiGroup := server.Group("/api")

	transactionGroup := apiGroup.Group("/client")
	{
		transactionGroup.POST("/change_balance", clientHandler.ChangeBalance)
	}

	_ = server.Run()
}
