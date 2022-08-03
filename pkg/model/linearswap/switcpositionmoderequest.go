package linearswap

type SwitchPositionModeRequest struct {
	MarginAccount string `json:"margin_account"`
	PositionMode  string `json:"position_mode"`
}
