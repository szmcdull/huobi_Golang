package market

import (
	"github.com/shopspring/decimal"
	"github.com/szmcdull/huobi_golang/pkg/model/base"
)

type SubscribeMarketByPriceResponse struct {
	base.WebSocketResponseBase
	Tick *MarketByPrice
	Data *MarketByPrice
}

type MarketByPrice struct {
	SeqNum     int64               `json:"seqNum"`
	PrevSeqNum int64               `json:"prevSeqNum"`
	Bids       [][]decimal.Decimal `json:"bids"`
	Asks       [][]decimal.Decimal `json:"asks"`
}
