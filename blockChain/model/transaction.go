package model

type Transaction struct{
	ID 	  string `json:"id_transaction"`
    Sender    string `json:"sender"`
    Receiver  string 	`json:"receiver"` 	
    Amount    float64 	`json:"amount"` 	
    Tax       float64 `json:"tax"` 	
    Timestamp string  `json:"timestamp"`
	Status    string  `json:"status"`	
    Signature string  `json:"signature"` 	 	
}