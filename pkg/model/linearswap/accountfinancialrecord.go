package linearswap

type AccountFinancialRecord struct {
	Id                int64   `json:"id"`
	Ts                int64   `json:"ts"`
	Asset             string  `json:"asset"`
	ContractCode      string  `json:"contract_code"`
	MarginAccount     string  `json:"margin_account"`
	FaceMarginAccount string  `json:"face_margin_account"`
	Type              int     `json:"type"`
	Amount            float64 `json:"amount"`
}
