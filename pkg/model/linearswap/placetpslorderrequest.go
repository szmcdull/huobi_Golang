package linearswap

type PlaceTPSLOrderRequest struct {
	ContractCode     string   `json:"contract_code"`
	Pair             string   `json:"pair"`
	ContractType     string   `json:"contract_type"`
	Direction        string   `json:"direction"`
	Volume           int64    `json:"volume"`
	TpTriggerPrice   *float64 `json:"tp_trigger_price"`
	TpOrderPrice     *float64 `json:"tp_order_price"`
	TpOrderPriceType *string  `json:"tp_order_price_type"`
	SlTriggerPrice   *float64 `json:"sl_trigger_price"`
	SlOrderPrice     *float64 `json:"sl_order_price"`
	SlOrderPriceType *string  `json:"sl_order_price_type"`
	ChannelCode      *string  `json:"channel_code"`
}
