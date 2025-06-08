package controller

import (
	"github.com/gin-gonic/gin"
	"go-api/usecase"
	"net/http"
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

func (tc *TransactionController) AccepetTransaction(ctx *gin.Context) {
	var req struct {
		ID          string `json:"id"`
		Assignature string `json:"assignature"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	transaction, err := tc.TransactionUsecase.GetTransactionByID(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if transaction == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transação não encontrada"})
		return
	}
	if transaction.Status != "PENDING" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Transação já foi aceita ou rejeitada"})
		return
	}
	_, err = tc.TransactionUsecase.AcceptTransaction(req.ID, req.Assignature)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Transação aceita"})
}

func (tc *TransactionController) GetTransactionByID(ctx *gin.Context) {
	id := ctx.Param("id")
	transaction, err := tc.TransactionUsecase.GetTransactionByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if transaction == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transação não encontrada"})
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}
