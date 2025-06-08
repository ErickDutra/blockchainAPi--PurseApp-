package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	//repository
	TransactionRepository := repository.NewTransactionRepository(dbConnection)
	BlockRepository := repository.NewBlockRepository(dbConnection)
	//usecase
	BlockUsecase := usecase.NewBlockUsecase(BlockRepository, TransactionRepository)
	TransactionUsecase := usecase.NewTransactionUsecase(TransactionRepository, &BlockUsecase)

	//controller
	transactionController := controller.NewTransactionController(TransactionUsecase)
	blockController := controller.NewBlockController(BlockUsecase)
	//routes
	server.GET("/transaction", transactionController.GetTransactions)
	server.POST("/transaction", transactionController.PostTransaction)
	server.POST("/transaction/accept/", transactionController.AccepetTransaction)
	server.GET("/transaction/:id", transactionController.GetTransactionByID)
	server.GET("/block", blockController.GetBlocks)
	server.Run(":8000")
}