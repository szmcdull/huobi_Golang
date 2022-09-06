package linearswap

type GetHistoryOrdersResponse struct {
	Status string                `json:"status"`
	Ts     int64                 `json:"ts"`
	Data   *GetHistoryOrdersData `json:"data"`
}

type GetHistoryOrdersData struct {
	Orders      []*HistoryOrder `json:"orders"`
	TotalPage   int             `json:"total_page"`
	CurrentPage int             `json:"current_page"`
	TotalSize   int             `json:"total_size"`
}

type GetHistoryOrdersV3Response struct {
	Status string                  `json:"status"`
	Ts     int64                   `json:"ts"`
	Data   *GetHistoryOrdersV3Data `json:"data"`
}

type GetHistoryOrdersV3Data struct {
	Orders      []*HistoryOrderV3 `json:"orders"`
	TotalPage   int               `json:"total_page"`
	CurrentPage int               `json:"current_page"`
	TotalSize   int               `json:"total_size"`
}
