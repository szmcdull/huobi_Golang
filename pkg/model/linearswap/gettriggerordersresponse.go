package linearswap

type GetTriggerOrdersResponse struct {
	Status  string                `json:"status"`
	ErrCode int                   `json:"err_code"`
	ErrMsg  string                `json:"err_msg"`
	Ts      int64                 `json:"ts"`
	Data    *GetTriggerOrdersData `json:"data"`
}

type GetTriggerOrdersData struct {
	Orders      []*TriggerOrder `json:"orders"`
	TotalPage   int             `json:"total_page"`
	CurrentPage int             `json:"current_page"`
	TotalSize   int             `json:"total_size"`
}
