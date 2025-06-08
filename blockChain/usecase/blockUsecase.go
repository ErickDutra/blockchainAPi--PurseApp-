package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/google/uuid"
	"go-api/model"
	"go-api/repository"
	"time"
)

type BlockUsecase struct {
	repository            repository.BlockRepository
	transactionRepository repository.TransactionRepository
}

func NewBlockUsecase(blockRepo repository.BlockRepository, transactionRepo repository.TransactionRepository) BlockUsecase {
	return BlockUsecase{
		repository:            blockRepo,
		transactionRepository: transactionRepo,
	}
}

func (tran *BlockUsecase) NewBlockGenesis(idTransaction string) (*model.Block, error) {
	transaction, err := tran.transactionRepository.GetTransactionByID(idTransaction)
	if err != nil {
		return nil, err
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	id := uuid.New().String()
	transactionData, err := json.Marshal(transaction)
	if err != nil {
		return nil, err
	}
	heshTransection := sha256.Sum256(transactionData)
	hashNovoTransection := hex.EncodeToString(heshTransection[:])

	blockAnterior, err := tran.repository.GetLastBlock()
	if err != nil {
		return nil, err
	}

	blockData, err := json.Marshal(blockAnterior)
	if err != nil {
		return nil, err
	}

	blockAnteriorHash := sha256.Sum256(blockData)
	hashAnterior := hex.EncodeToString(blockAnteriorHash[:])

	block := model.Block{
		Index:           id,
		HashTransaction: hashNovoTransection,
		HashAnterior:    hashAnterior,
		Timestamp:       timestamp,
	}
	err = tran.repository.PostBlock(block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

// func CreateBlock( index string,data string, hashTransaction string, hashAnterior string ) *string {
// 	timestamp := time.Now().Format("2006-01-02 15:04:05")
// 	block := Block{index, data, timestamp, hashTransaction, hashAnterior}

// 	blockData, err := json.Marshal(block)
// 	if err != nil {
// 		errorMessage := "Falha ao converter o bloco em JSON"
// 		return &errorMessage
// 	}

//		heshBlock := sha256.Sum256(blockData)
//		hashNovoBlock := hex.EncodeToString(heshBlock[:])
//		return &hashNovoBlock
//	}
func (tran *BlockUsecase) GetAllBlocks() ([]model.Block, error) {
	blocks, err := tran.repository.GetAllBlocks()
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

func (tran *BlockUsecase) GetLastBlock() (*model.Block, error) {
	block, err := tran.repository.GetLastBlock()
	if err != nil {
		return nil, err
	}
	return block, nil
}
