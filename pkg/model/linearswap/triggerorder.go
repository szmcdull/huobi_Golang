package linearswap

type TriggerOrder struct {
	ContractType    string   `json:"contract_type"`
	BusinessType    string   `json:"business_type"`
	Pair            string   `json:"pair"`
	Symbol          string   `json:"symbol"`
	ContractCode    string   `json:"contract_code"`
	TriggerType     string   `json:"trigger_type"`
	Volume          float64  `json:"volume"`
	OrderType       int      `json:"order_type"`
	Direction       string   `json:"direction"`
	Offset          string   `json:"offset"`
	LeverRate       int      `json:"lever_rate"`
	OrderId         int64    `json:"order_id"`
	OrderIdStr      string   `json:"order_id_str"`
	RelationOrderId string   `json:"relation_order_id"`
	OrderSource     string   `json:"order_source"`
	TriggerPrice    float64  `json:"trigger_price"`
	TriggeredPrice  *float64 `json:"triggered_price"`
	OrderPrice      float64  `json:"order_price"`
	CreatedAt       int64    `json:"created_at"`
	TriggeredAt     *int64   `json:"triggered_at"`
	OrderPriceType  string   `json:"order_price_type"`
	Status          int      `json:"status"`
	MarginMode      string   `json:"margin_mode"`
	MarginAccount   string   `json:"margin_account"`
	ReduceOnly      int      `json:"reduce_only"`
}
