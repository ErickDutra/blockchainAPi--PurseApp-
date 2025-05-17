package model


type Status struct {
    Status string `json:"status"`
}


const (
    StatusPending   = "PENDING"
    StatusConfirmed = "CONFIRMED"
    StatusRejected    = "REJECTED"
)