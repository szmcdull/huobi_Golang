package linearswap

type CrossLinearSwapUserAccountInfoResponse struct {
	Status string                           `json:"status"`
	Data   []CrossLinearSwapUserAccountInfo `json:"data"`
}

type FuturesContractDetail struct {
	Symbol           string   `json:"symbol"`
	ContractCode     string   `json:"contract_code"`
	MarginPosition   float64  `json:"margin_position"`
	MarginFrozen     float64  `json:"margin_frozen"`
	MarginAvailable  float64  `json:"margin_available"`
	ProfitUnreal     float64  `json:"profit_unreal"`
	LiquidationPrice *float64 `json:"liquidation_price"`
	LeverRate        int      `json:"lever_rate"`
	AdjustFactor     float64  `json:"adjust_factor"`
	ContractType     string   `json:"contract_type"`
	Pair             string   `json:"pair"`
	BusinessType     string   `json:"business_type"`
	TradePartition   string   `json:"trade_partition"`
}

type ContractDetail struct {
	Symbol           string   `json:"symbol"`
	ContractCode     string   `json:"contract_code"`
	MarginPosition   float64  `json:"margin_position"`
	MarginFrozen     float64  `json:"margin_frozen"`
	MarginAvailable  float64  `json:"margin_available"`
	ProfitUnreal     int      `json:"profit_unreal"`
	LiquidationPrice *float64 `json:"liquidation_price"`
	LeverRate        int      `json:"lever_rate"`
	AdjustFactor     float64  `json:"adjust_factor"`
	ContractType     string   `json:"contract_type"`
	Pair             string   `json:"pair"`
	BusinessType     string   `json:"business_type"`
	TradePartition   string   `json:"trade_partition"`
}

type CrossLinearSwapUserAccountInfo struct {
	FuturesContractDetail []*FuturesContractDetail `json:"futures_contract_detail"`
	MarginMode            string                   `json:"margin_mode"`
	MarginAccount         string                   `json:"margin_account"`
	MarginAsset           string                   `json:"margin_asset"`
	MarginBalance         float64                  `json:"margin_balance"`
	MarginStatic          float64                  `json:"margin_static"`
	MarginPosition        float64                  `json:"margin_position"`
	MarginFrozen          float64                  `json:"margin_frozen"`
	ProfitReal            float64                  `json:"profit_real"`
	ProfitUnreal          float64                  `json:"profit_unreal"`
	WithdrawAvailable     float64                  `json:"withdraw_available"`
	RiskRate              *float64                 `json:"risk_rate"`
	ContractDetail        []*ContractDetail        `json:"contract_detail"`
}
