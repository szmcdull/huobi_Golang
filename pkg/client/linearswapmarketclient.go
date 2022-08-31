package client

import (
	"encoding/json"
	"errors"
	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/market"
	"strconv"
)

// Responsible to get market information
type LinearSwapMarketClient struct {
	publicUrlBuilder *requestbuilder.PublicUrlBuilder
}

// Initializer
func (p *LinearSwapMarketClient) Init(host string) *LinearSwapMarketClient {
	p.publicUrlBuilder = new(requestbuilder.PublicUrlBuilder).Init(host)
	return p
}

// Retrieves the latest trade with its price, volume, and direction.
func (client *LinearSwapMarketClient) GetLatestTrade(symbol string) (*market.TradeTick, error) {

	request := new(model.GetRequest).Init()
	request.AddParam("contract_code", symbol)

	url := client.publicUrlBuilder.Build("/linear-swap-ex/market/trade", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetLatestTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Tick != nil {

		return result.Tick, nil
	}

	return nil, errors.New(getResp)
}

// Retrieves the most recent trades with their price, volume, and direction.
func (client *LinearSwapMarketClient) GetHistoricalTrade(symbol string, optionalRequest market.GetHistoricalTradeOptionalRequest) ([]market.TradeTick, error) {

	request := new(model.GetRequest).Init()
	request.AddParam("contract_code", symbol)

	if optionalRequest.Size != 0 {
		request.AddParam("size", strconv.Itoa(optionalRequest.Size))
	}

	url := client.publicUrlBuilder.Build("/linear-swap-ex/market/history/trade", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetHistoricalTradeResponse{}

	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}
