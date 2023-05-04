package linearswap

type GetRelatedToPositionTPSLOrderRequest struct {
	ContractCode string `json:"contract_code"`
	Pair         string `json:"pair"`
	OrderId      uint64 `json:"order_id"`
}
