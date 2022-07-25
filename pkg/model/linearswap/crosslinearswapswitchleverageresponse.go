package linearswap

type CrossLinearSwapSwitchLeverageResponse struct {
	Status  string                                     `json:"status"`
	Ts      int64                                      `json:"ts"`
	ErrCode int                                        `json:"err-code"`
	ErrMsg  string                                     `json:"err-msg"`
	Data    *CrossLinearSwapSwitchLeverageResponseData `json:"data"`
}

type CrossLinearSwapSwitchLeverageResponseData struct {
	ContractCode string `json:"contract_code"`
	MarginMode   string `json:"margin_mode"`
	LeverRate    int    `json:"lever_rate"`
	ContractType string `json:"contract_type"`
	Pair         string `json:"pair"`
	BusinessType string `json:"business_type"`
}
