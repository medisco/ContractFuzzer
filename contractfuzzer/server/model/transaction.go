package model

type TransactionCreateRequest struct {
	TaskId         string `json:"taskId"`
	BlockchainHash string `json:"blockchainHash"`
}

type TransactionCreateResponse struct {
	TransactionId string `json:"transactionId"`
}
