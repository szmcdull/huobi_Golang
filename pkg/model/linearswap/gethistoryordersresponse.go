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
