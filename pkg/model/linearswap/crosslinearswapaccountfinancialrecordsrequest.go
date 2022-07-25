package linearswap

type CrossLinearSwapAccountFinancialRecordsRequest struct {
	MarginAccount string `json:"margin_account"`
	ContractCode  string `json:"contract_code"`
	Type          string `json:"type"`
	CreateDate    int    `json:"create_date"`
	PageIndex     int    `json:"page_index"`
	PageSize      int    `json:"page_size"`
}
