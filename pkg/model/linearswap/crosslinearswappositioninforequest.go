package linearswap

type CrossLinearSwapPositionInfoRequest struct {
	ContractCode string `json:"contract_code,omitempty"`
	Pair         string `json:"pair,omitempty"`
	ContractType string `json:"contract_type,omitempty"`
}
