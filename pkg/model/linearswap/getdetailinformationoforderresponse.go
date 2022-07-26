package linearswap

type GetDetailInformationOfOrderResponse struct {
	Status string                    `json:"status"`
	Ts     int64                     `json:"ts"`
	Data   *DetailInformationOfOrder `json:"data"`
}

type DetailInformationOfOrder struct {
	ContractType    string   `json:"contract_type"`
	Pair            string   `json:"pair"`
	BusinessType    string   `json:"business_type"`
	Symbol          string   `json:"symbol"`
	ContractCode    string   `json:"contract_code"`
	InstrumentPrice int      `json:"instrument_price"`
	FinalInterest   int      `json:"final_interest"`
	AdjustValue     int      `json:"adjust_value"`
	LeverRate       int      `json:"lever_rate"`
	Direction       string   `json:"direction"`
	Offset          string   `json:"offset"`
	Volume          float64  `json:"volume"`
	Price           float64  `json:"price"`
	CreatedAt       int64    `json:"created_at"`
	CanceledAt      int      `json:"canceled_at"`
	OrderSource     string   `json:"order_source"`
	OrderPriceType  string   `json:"order_price_type"`
	MarginFrozen    float64  `json:"margin_frozen"`
	Profit          float64  `json:"profit"`
	TotalPage       int      `json:"total_page"`
	CurrentPage     int      `json:"current_page"`
	TotalSize       int      `json:"total_size"`
	LiquidationType string   `json:"liquidation_type"`
	FeeAsset        string   `json:"fee_asset"`
	Fee             float64  `json:"fee"`
	OrderId         int64    `json:"order_id"`
	OrderIdStr      string   `json:"order_id_str"`
	ClientOrderId   *string  `json:"client_order_id"`
	OrderType       string   `json:"order_type"`
	Status          int      `json:"status"`
	TradeAvgPrice   float64  `json:"trade_avg_price"`
	TradeTurnover   float64  `json:"trade_turnover"`
	TradeVolume     float64  `json:"trade_volume"`
	MarginAsset     string   `json:"margin_asset"`
	MarginAccount   string   `json:"margin_account"`
	MarginMode      string   `json:"margin_mode"`
	IsTpsl          int      `json:"is_tpsl"`
	RealProfit      int      `json:"real_profit"`
	ReduceOnly      int      `json:"reduce_only"`
	Trades          []*Trade `json:"trades"`
}

type Trade struct {
	TradeId       int     `json:"trade_id"`
	TradePrice    float64 `json:"trade_price"`
	TradeVolume   float64 `json:"trade_volume"`
	TradeTurnover float64 `json:"trade_turnover"`
	TradeFee      float64 `json:"trade_fee"`
	CreatedAt     int64   `json:"created_at"`
	Role          string  `json:"role"`
	FeeAsset      string  `json:"fee_asset"`
	RealProfit    float64 `json:"real_profit"`
	Profit        float64 `json:"profit"`
	Id            string  `json:"id"`
}
