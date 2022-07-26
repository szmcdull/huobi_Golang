package linearswap

type GetInformationOfOrderResponse struct {
	Status string   `json:"status"`
	Ts     int64    `json:"ts"`
	Data   []*Order `json:"data"`
}
