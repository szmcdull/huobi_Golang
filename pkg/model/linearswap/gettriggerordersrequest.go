package linearswap

type GetTriggerOrdersRequest struct {
	ContractCode string `json:"contract_code"`
	Pair         string `json:"pair"`
	PageIndex    int    `json:"page_index"`
	PageSize     int    `json:"page_size"`
	TradeType    int    `json:"trade_type"`
}
