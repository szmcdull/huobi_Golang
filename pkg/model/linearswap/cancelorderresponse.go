package linearswap

type CancelOrderResponse struct {
	Status string             `json:"status"`
	Data   *CancelOrderResult `json:"data"`
	Ts     int64              `json:"ts"`
}

type CancelOrderResult struct {
	Errors    []*CancelOrderError `json:"errors"`
	Successes string              `json:"successes"`
}

type CancelOrderError struct {
	OrderId string `json:"order_id"`
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}
