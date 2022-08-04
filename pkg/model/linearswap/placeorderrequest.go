package linearswap

type PlaceOrderRequest struct {
	ContractCode     string   `json:"contract_code"`
	Pair             string   `json:"pair"`
	ContractType     string   `json:"contract_type"`
	ReduceOnly       int      `json:"reduce_only"`
	ClientOrderId    *int64   `json:"client_order_id"`
	Price            *float64 `json:"price"`
	Volume           int64    `json:"volume"`
	Direction        string   `json:"direction"`
	Offset           *string  `json:"offset"`
	LeverRate        int      `json:"lever_rate"`
	OrderPriceType   string   `json:"order_price_type"`
	TpTriggerPrice   *float64 `json:"tp_trigger_price"`
	TpOrderPrice     *float64 `json:"tp_order_price"`
	TpOrderPriceType *string  `json:"tp_order_price_type"`
	SlTriggerPrice   *float64 `json:"sl_trigger_price"`
	SlOrderPrice     *float64 `json:"sl_order_price"`
	SlOrderPriceType *string  `json:"sl_order_price_type"`
	ChannelCode      *string  `json:"channel_code"`
}
