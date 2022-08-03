package account

type TransferAccountRequest struct {
	FromUser        int64  `json:"from-user"`
	FromAccountType string `json:"from-account-type"`
	FromAccount     int64  `json:"from-account"`
	ToUser          int64  `json:"to-user"`
	ToAccountType   string `json:"to-account-type"`
	ToAccount       int64  `json:"to-account"`
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
}

type TransferAccountRequestV2 struct {
	From          string  `json:"from"`
	To            string  `json:"to"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	MarginAccount string  `json:"margin-account"`
}
