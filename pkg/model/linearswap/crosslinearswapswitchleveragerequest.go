package linearswap

type CrossLinearSwapSwitchLeverageRequest struct {
	ContractCode *string `json:"contract_code"`
	Pair         *string `json:"pair"`
	ContractType *string `json:"contract_type"`
	LeverRate    int     `json:"lever_rate"`
}
