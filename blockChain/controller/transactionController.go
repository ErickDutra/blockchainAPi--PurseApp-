package controller

import (
	"go-api/usecase"
	"net/http"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionUsecase usecase.TransactionUsecase

}

func NewTransactionController(usecase usecase.TransactionUsecase) TransactionController {
	return TransactionController{
		TransactionUsecase: usecase,
	}
}

func (tc *TransactionController) PostTransaction(ctx *gin.Context) {
    var req struct {
        Sender   string  `json:"sender"`
        Receiver string  `json:"receiver"`
        Amount   float64 `json:"amount"`
        Tax      float64 `json:"tax"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
        return
    }

    transaction, err := tc.TransactionUsecase.NewTransaction(req.Sender, req.Receiver, req.Amount, req.Tax)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, transaction)
}

func (tc *TransactionController) GetTransactions(ctx *gin.Context) {
	transactions, err := tc.TransactionUsecase.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transactions)

}