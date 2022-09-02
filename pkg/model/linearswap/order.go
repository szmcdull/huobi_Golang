package linearswap

type Order struct {
	BusinessType    string   `json:"business_type"`
	ContractType    string   `json:"contract_type"`
	Pair            string   `json:"pair"`
	Symbol          string   `json:"symbol"`
	ContractCode    string   `json:"contract_code"`
	Volume          float64  `json:"volume"`
	Price           float64  `json:"price"`
	OrderPriceType  string   `json:"order_price_type"`
	OrderType       int      `json:"order_type"`
	Direction       string   `json:"direction"`
	Offset          string   `json:"offset"`
	LeverRate       int      `json:"lever_rate"`
	OrderId         int64    `json:"order_id"`
	ClientOrderId   *int64   `json:"client_order_id"`
	CreatedAt       int64    `json:"created_at"`
	TradeVolume     float64  `json:"trade_volume"`
	TradeTurnover   float64  `json:"trade_turnover"`
	Fee             float64  `json:"fee"`
	TradeAvgPrice   *float64 `json:"trade_avg_price"`
	MarginFrozen    float64  `json:"margin_frozen"`
	Profit          float64  `json:"profit"`
	Status          int      `json:"status"`
	OrderSource     string   `json:"order_source"`
	OrderIdStr      string   `json:"order_id_str"`
	FeeAsset        string   `json:"fee_asset"`
	LiquidationType string   `json:"liquidation_type"`
	CanceledAt      int      `json:"canceled_at"`
	MarginAsset     string   `json:"margin_asset"`
	MarginAccount   string   `json:"margin_account"`
	MarginMode      string   `json:"margin_mode"`
	IsTpsl          int      `json:"is_tpsl"`
	RealProfit      int      `json:"real_profit"`
	ReduceOnly      int      `json:"reduce_only"`
}
