package linearswap

type GetHistoryOrdersViaMultipleFieldsResponse struct {
	Status string                                 `json:"status"`
	Ts     int64                                  `json:"ts"`
	Data   *GetHistoryOrdersViaMultipleFieldsData `json:"data"`
}

type GetHistoryOrdersViaMultipleFieldsData struct {
	Orders     []*HistoryOrderMultipleFields `json:"orders"`
	RemainSize int                           `json:"remain_size"`
	NextId     *int64                        `json:"next_id"`
}

type GetHistoryOrdersViaMultipleFieldsV3Response struct {
	Code int                             `json:"code"`
	Msg  string                          `json:"msg"`
	Ts   int64                           `json:"ts"`
	Data []*HistoryOrderMultipleFieldsV3 `json:"data"`
}
