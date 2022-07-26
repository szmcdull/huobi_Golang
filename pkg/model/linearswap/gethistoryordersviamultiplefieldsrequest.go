package linearswap

type GetHistoryOrdersViaMultipleFieldsRequest struct {
	ContractCode   string `json:"contract_code"`
	Pair           string `json:"pair"`
	TradeType      int    `json:"trade_type"`
	Type           int    `json:"type"`
	Status         string `json:"status"`
	OrderPriceType string `json:"order_price_type"`
	StartTime      int64  `json:"start_time"`
	EndTime        int64  `json:"end_time"`
	FromId         int64  `json:"from_id"`
	Size           int    `json:"size"`
	Direct         string `json:"direct"`
}
