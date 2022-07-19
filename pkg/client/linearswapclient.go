package client

import (
	"encoding/json"
	"errors"
	"github.com/huobirdcenter/huobi_golang/internal"
	"github.com/huobirdcenter/huobi_golang/internal/requestbuilder"
	"github.com/huobirdcenter/huobi_golang/pkg/model"
	"github.com/huobirdcenter/huobi_golang/pkg/model/linearswap"
)

// Responsible to operate linear swap
type LinearSwapClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *LinearSwapClient) Init(accessKey string, secretKey string, host string) *LinearSwapClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

// Returns linear swap positions info.
func (p *LinearSwapClient) QueryUserPositionInfo(request *linearswap.CrossLinearSwapPositionInfoRequest) ([]linearswap.CrossLinearSwapPositionInfo, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_position_info", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}
	result := linearswap.CrossLinearSwapPositionInfoResponse{}
	if jsonErr = json.Unmarshal([]byte(postResp), &result); jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(postResp)
}

// Returns user's linear swap account info.
func (p *LinearSwapClient) QueryUserAccountInfo(request *linearswap.CrossLinearSwapUserAccountInfoRequest) ([]linearswap.CrossLinearSwapUserAccountInfo, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}
	url := p.privateUrlBuilder.Build("POST", "/linear-swap-api/v1/swap_cross_account_info", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}
	result := linearswap.CrossLinearSwapUserAccountInfoResponse{}
	if jsonErr = json.Unmarshal([]byte(postResp), &result); jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(postResp)
}
