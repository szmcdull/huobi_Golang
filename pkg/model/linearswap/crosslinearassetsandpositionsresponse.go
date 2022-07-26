package linearswap

type CrossLinearAssetsAndPositionsResponse struct {
	Status string                             `json:"status"`
	Ts     int64                              `json:"ts"`
	Data   *CrossLinearAssetsAndPositionsData `json:"data"`
}
type CrossLinearAssetsAndPositionsData struct {
	Positions             []*CrossLinearAssetsAndPositionsPosition `json:"positions"`
	FuturesContractDetail []*ContractDetail                        `json:"futures_contract_detail"`
	MarginMode            string                                   `json:"margin_mode"`
	MarginAccount         string                                   `json:"margin_account"`
	MarginAsset           string                                   `json:"margin_asset"`
	MarginBalance         float64                                  `json:"margin_balance"`
	MarginStatic          float64                                  `json:"margin_static"`
	MarginPosition        float64                                  `json:"margin_position"`
	MarginFrozen          float64                                  `json:"margin_frozen"`
	ProfitReal            float64                                  `json:"profit_real"`
	ProfitUnreal          float64                                  `json:"profit_unreal"`
	WithdrawAvailable     float64                                  `json:"withdraw_available"`
	RiskRate              float64                                  `json:"risk_rate"`
	PositionMode          string                                   `json:"position_mode"`
	ContractDetail        []*ContractDetail                        `json:"contract_detail"`
}

type CrossLinearAssetsAndPositionsPosition struct {
	Symbol         string  `json:"symbol"`
	ContractCode   string  `json:"contract_code"`
	Volume         float64 `json:"volume"`
	Available      float64 `json:"available"`
	Frozen         float64 `json:"frozen"`
	CostOpen       float64 `json:"cost_open"`
	CostHold       float64 `json:"cost_hold"`
	ProfitUnreal   float64 `json:"profit_unreal"`
	ProfitRate     float64 `json:"profit_rate"`
	LeverRate      int     `json:"lever_rate"`
	PositionMargin float64 `json:"position_margin"`
	Direction      string  `json:"direction"`
	Profit         float64 `json:"profit"`
	LastPrice      float64 `json:"last_price"`
	MarginAsset    string  `json:"margin_asset"`
	MarginMode     string  `json:"margin_mode"`
	MarginAccount  string  `json:"margin_account"`
	ContractType   string  `json:"contract_type"`
	Pair           string  `json:"pair"`
	BusinessType   string  `json:"business_type"`
	PositionMode   string  `json:"position_mode"`
}

type CrossLinearAssetsAndPositionsContractDetail struct {
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
}
