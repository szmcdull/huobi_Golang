package linearswap

type SwapTradingFee struct {
	Symbol        string `json:"symbol"`
	ContractCode  string `json:"contract_code"`
	OpenMakerFee  string `json:"open_maker_fee"`
	OpenTakerFee  string `json:"open_taker_fee"`
	CloseMakerFee string `json:"close_maker_fee"`
	CloseTakerFee string `json:"close_taker_fee"`
	FeeAsset      string `json:"fee_asset"`
	DeliveryFee   string `json:"delivery_fee"`
	BusinessType  string `json:"business_type"`
	ContractType  string `json:"contract_type"`
	Pair          string `json:"pair"`
}

type SwapTradingFeeRequest struct {
	ContractCode *string `json:"contract_code"`
	Pair         *string `json:"pair"`
	ContractType *string `json:"contract_type"`
	BusinessType *string `json:"business_type"`
}

type SwapTradingFeeResponse struct {
	Status string            `json:"status"`
	Ts     int64             `json:"ts"`
	Data   []*SwapTradingFee `json:"data"`
}
