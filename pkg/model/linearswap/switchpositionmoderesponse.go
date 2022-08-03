package linearswap

type SwitchPositionModeResponse struct {
	Status string          `json:"status"`
	Data   []*PositionMode `json:"data"`
	Ts     int64           `json:"ts"`
}

type PositionMode struct {
	MarginAccount string `json:"margin_account"`
	PositionMode  string `json:"position_mode"`
}
