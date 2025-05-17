package main; 
// import (
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"encoding/json"
// 	"time"
// )



// func CreateBlock( index string,data string, hashTransaction string, hashAnterior string ) *string {
// 	timestamp := time.Now().Format("2006-01-02 15:04:05")
// 	block := Block{index, data, timestamp, hashTransaction, hashAnterior}

// 	blockData, err := json.Marshal(block)
// 	if err != nil {
// 		errorMessage := "Falha ao converter o bloco em JSON"
// 		return &errorMessage
// 	}
	
// 	heshBlock := sha256.Sum256(blockData)
// 	hashNovoBlock := hex.EncodeToString(heshBlock[:])
// 	return &hashNovoBlock
// }




// func NewTransaction(sender string, receiver string, amount float64, tax float64,status string, signature string)*string {
// 	timestamp := time.Now().Format("2006-01-02 15:04:05")
// 	transaction := Transaction{sender, receiver, amount, tax, timestamp, signature}
// 	data, err := json.Marshal(transaction)
// 	if err != nil {
// 		 errorMessage := "Falha ao converter a transação em JSON"
//         return &errorMessage
// 	}
// 	heshTransaction := sha256.Sum256(data)	
// 	hashNewTransaction:= hex.EncodeToString(heshTransaction[:])
// 	return &hashNewTransaction 
// }

