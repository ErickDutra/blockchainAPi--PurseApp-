package usecase

import (
	"go-api/model"
	"go-api/repository"
	"time"

	"github.com/google/uuid"
)


type TransactionUsecase struct{
	repository repository.TransactionRepository
}

func NewTransactionUsecase(repo repository.TransactionRepository) TransactionUsecase {
	return TransactionUsecase{
		repository: repo,
	}
}

func (tran *TransactionUsecase) NewTransaction(sender string, receiver string, amount float64, tax float64) (*model.Transaction, error) {
    timestamp := time.Now().Format("2006-01-02 15:04:05")
    status := model.StatusPending
    signature := ""
    id := uuid.New().String()
    transaction := model.Transaction{
        ID:        id,
        Sender:    sender,
        Receiver:  receiver,
        Amount:    amount,
        Tax:       tax,
        Timestamp: timestamp,
        Status:    status,
        Signature: signature,
    }
    err := tran.repository.PostTransaction(transaction)
    if err != nil {
        return nil, err
    }
    return &transaction, nil
}

func (tran *TransactionUsecase) GetAllTransactions() ([]model.Transaction, error) {
	transactions, err := tran.repository.GetAllTransactions()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}