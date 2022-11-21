package linearswap

type GetLiquidationOrdersRequest struct {
	Contract  string  `json:"contract"`
	TradeType int     `json:"trade_type"`
	Pair      *string `json:"pair"`
	StartTime *int64  `json:"start_time"`
	EndTime   *int64  `json:"end_time"`
	Direct    *string `json:"direct"`
	FromId    *int64  `json:"from_id"`
}

type GetLiquidationOrdersResponse struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Ts   int64               `json:"ts"`
	Data []*LiquidationOrder `json:"data"`
}

type LiquidationOrder struct {
	QueryID       int     `json:"query_id"`
	ContractCode  string  `json:"contract_code"`
	Symbol        string  `json:"symbol"`
	Direction     string  `json:"direction"`
	Offset        string  `json:"offset"`
	Volume        float64 `json:"volume"`
	Price         float64 `json:"price"`
	CreatedAt     int64   `json:"created_at"`
	Amount        float64 `json:"amount"`
	TradeTurnover float64 `json:"trade_turnover"`
	BusinessType  string  `json:"business_type"`
	Pair          string  `json:"pair"`
}
