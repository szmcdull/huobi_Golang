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
	Size          int32   `json:"size"`
	Direct        *string `json:"direct"`
}

type CrossLinearSwapAccountFinancialRecordsV3Request struct {
	Contract  *string `json:"contract"`
	MarAcct   string  `json:"mar_acct"`
	Type      *string `json:"type"`
	StartTime *int64  `json:"start_time"`
	EndTime   *int64  `json:"end_time"`
	Direct    *string `json:"direct"`
	FromId    *int    `json:"from_id"`
}

type QueryAccountFinancialRecordsViaMultipleFieldsV3Request struct {
	Contract  string `json:"contract"`
	MarAcct   string `json:"mar_acct"`
	Type      string `json:"type"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
	Direct    string `json:"direct"`
	FromId    int64  `json:"from_id"`
}
