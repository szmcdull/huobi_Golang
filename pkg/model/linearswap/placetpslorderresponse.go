package linearswap

type PlaceTPSLOrderResponse struct {
	Status  string                      `json:"status"`
	ErrCode int                         `json:"err_code"`
	ErrMsg  string                      `json:"err_msg"`
	Data    *PlaceTPSLOrderResponseData `json:"data"`
	Ts      int64                       `json:"ts"`
}

type PlaceTPSLOrderResponseData struct {
	TpOrder *TPSLOrderIdentifier `json:"tp_order"`
	SlOrder *TPSLOrderIdentifier `json:"sl_order"`
}

type TPSLOrderIdentifier struct {
	OrderId    int64  `json:"order_id"`
	OrderIdStr string `json:"order_id_str"`
}
