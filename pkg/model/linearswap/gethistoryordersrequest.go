package linearswap

type GetHistoryOrdersRequest struct {
	ContractCode string `json:"contract_code"`
	Pair         string `json:"pair"`
	TradeType    int    `json:"trade_type"`
	Type         int    `json:"type"`
	Status       string `json:"status"`
	CreateDate   int    `json:"create_date"`
	PageIndex    int    `json:"page_index"`
	PageSize     int    `json:"page_size"`
	SortBy       string `json:"sort_by"`
}
