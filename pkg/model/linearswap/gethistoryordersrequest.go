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

type GetHistoryOrdersV3Request struct {
	Contract  string  `json:"contract"`
	TradeType int     `json:"trade_type"`
	Status    int     `json:"status"`
	Type      int     `json:"type"`
	Pair      *string `json:"pair"`
	StartTime *int64  `json:"start_time"`
	EndTime   *int64  `json:"end_time"`
	Direct    *string `json:"direct"`
	FromId    *int    `json:"from_id"`
}
