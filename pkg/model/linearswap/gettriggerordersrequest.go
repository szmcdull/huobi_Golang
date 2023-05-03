package linearswap

type GetTriggerOrdersRequest struct {
	ContractCode string `json:"contract_code"`
	Pair         string `json:"pair"`
	PageIndex    int    `json:"page_index"`
	PageSize     int    `json:"page_size"`
	TradeType    int    `json:"trade_type"`
}

type GetHistoryTriggerOrdersRequest struct {
	ContractCode *string `json:"contract_code"`
	Pair         *string `json:"pair"`
	TradeType    int     `json:"trade_type"`
	Status       string  `json:"status"`
	CreatedDate  int     `json:"create_date"`
	PageIndex    *int    `json:"page_index"`
}
