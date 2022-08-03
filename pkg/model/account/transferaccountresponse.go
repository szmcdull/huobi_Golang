package account

type TransferAccountResponse struct {
	Status string              `json:"status"`
	Data   TransferAccountData `json:"data"`
}

type TransferAccountData struct {
	TransactId   int64 `json:"transact-id"`
	TransactTime int64 `json:"transact-time"`
}

type TransferAccountResponseV2 struct {
	Success bool   `json:"success"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    int64  `json:"data"`
}
