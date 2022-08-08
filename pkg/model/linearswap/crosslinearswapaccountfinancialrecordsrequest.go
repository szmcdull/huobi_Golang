package linearswap

type CrossLinearSwapAccountFinancialRecordsRequest struct {
	MarginAccount string `json:"margin_account"`
	ContractCode  string `json:"contract_code"`
	Type          string `json:"type"`
	CreateDate    int    `json:"create_date"`
	PageIndex     int    `json:"page_index"`
	PageSize      int    `json:"page_size"`
}

type QueryAccountFinancialRecordsViaMultipleFieldsRequest struct {
	MarginAccount string  `json:"margin_account"`
	ContractCode  string  `json:"contract_code"`
	Type          string  `json:"type"`
	StartTime     *int64  `json:"start_time"`
	EndTime       *int64  `json:"end_time"`
	FromId        *int64  `json:"from_id"`
	Direct        *string `json:"direct"`
}
