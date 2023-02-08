package linearswap

type CancelTriggerOrderRequest struct {
	OrderId      string `json:"order_id"`
	ContractCode string `json:"contract_code"`
	Pair         string `json:"pair"`
	ContractType string `json:"contract_type"`
}
