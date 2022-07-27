package linearswap

type GetDetailInformationOfOrderRequest struct {
	ContractCode string `json:"contract_code"`
	Pair         string `json:"pair"`
	OrderId      int64  `json:"order_id"`
	CreatedAt    *int64 `json:"created_at"`
	OrderType    *int   `json:"order_type"`
	PageIndex    int    `json:"page_index"`
	PageSize     int    `json:"page_size"`
}
