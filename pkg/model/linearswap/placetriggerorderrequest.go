package linearswap

type PlaceTriggerOrderRequest struct {
	ContractCode   string   `json:"contract_code"`
	Pair           string   `json:"pair"`
	ContractType   string   `json:"contract_type"`
	ReduceOnly     int      `json:"reduce_only"`
	TriggerType    string   `json:"trigger_type"`
	TriggerPrice   float64  `json:"trigger_price"`
	OrderPrice     *float64 `json:"order_price"`
	OrderPriceType *string  `json:"order_price_type"`
	Volume         int64    `json:"volume"`
	Direction      string   `json:"direction"`
	Offset         *string  `json:"offset"`
	LeverRate      int      `json:"lever_rate"`
	ChannelCode    *string  `json:"channel_code"`
}
