package linearswap

type CrossLinearSwapAccountFinancialRecord struct {
	Id                int64   `json:"id"`
	Ts                int64   `json:"ts"`
	Asset             string  `json:"asset"`
	ContractCode      string  `json:"contract_code"`
	MarginAccount     string  `json:"margin_account"`
	FaceMarginAccount string  `json:"face_margin_account"`
	Type              int     `json:"type"`
	Amount            float64 `json:"amount"`
}

type CrossLinearSwapAccountFinancialRecordV3 struct {
	QueryId           int64   `json:"query_id"`
	Id                int64   `json:"id"`
	Type              int     `json:"type"`
	Amount            float64 `json:"amount"`
	Ts                int64   `json:"ts"`
	ContractCode      string  `json:"contract_code"`
	Asset             string  `json:"asset"`
	MarginAccount     string  `json:"margin_account"`
	FaceMarginAccount string  `json:"face_margin_account"`
}
