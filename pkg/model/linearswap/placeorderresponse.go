package linearswap

type PlaceOrderResponse struct {
	Status string          `json:"status"`
	Ts     int64           `json:"ts"`
	Data   *PlaceOrderData `json:"data"`
}

type PlaceOrderData struct {
	OrderId       int64  `json:"order_id"`
	OrderIdStr    string `json:"order_id_str"`
	ClientOrderId *int64 `json:"client_order_id"`
}
