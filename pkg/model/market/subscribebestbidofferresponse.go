package market

import (
	"github.com/shopspring/decimal"
	"github.com/szmcdull/huobi_golang/pkg/model/base"
)

type SubscribeBestBidOfferResponse struct {
	base.WebSocketResponseBase
	Tick *struct {
		QuoteTime int64           `json:"quoteTime"`
		Symbol    string          `json:"symbol"`
		Bid       decimal.Decimal `json:"bid"`
		BidSize   decimal.Decimal `json:"bidSize"`
		Ask       decimal.Decimal `json:"ask"`
		AskSize   decimal.Decimal `json:"askSize"`
	}
}
