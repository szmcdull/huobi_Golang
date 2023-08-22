package subuser

import "github.com/shopspring/decimal"

type SubUserTransferRequest struct {
	SubUid   int64           `json:"sub-uid"`
	Currency string          `json:"currency"`
	Amount   decimal.Decimal `json:"amount"`
	Type     string          `json:"type"`
}

type SubUserTransferResponse struct {
	Status  string `json:"status"`
	ErrCode string `json:"err-code"`
	ErrMsg  string `json:"err-msg"`
	Data    *int64 `json:"data"`
}
