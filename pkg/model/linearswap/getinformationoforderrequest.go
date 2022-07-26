package linearswap

type GetInformationOfOrderRequest struct {
	OrderId       *string `json:"order_id"`
	ClientOrderId *string `json:"client_order_id"`
	ContractCode  *string `json:"contract_code"`
	Pair          *string `json:"pair"`
}
