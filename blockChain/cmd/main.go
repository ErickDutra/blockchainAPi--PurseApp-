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
	//usecase
	TransactionUsecase := usecase.NewTransactionUsecase(TransactionRepository)
	//controller
	transactionController := controller.NewTransactionController(TransactionUsecase)
	//routes
	server.GET("/transaction", transactionController.GetTransactions)
	server.POST("/transaction", transactionController.PostTransaction)

	server.Run(":8000")
}