package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/szmcdull/huobi_golang/internal"
	"github.com/szmcdull/huobi_golang/internal/requestbuilder"
	"github.com/szmcdull/huobi_golang/pkg/model"
	"github.com/szmcdull/huobi_golang/pkg/model/account"
	"github.com/szmcdull/huobi_golang/pkg/model/subuser"
	"github.com/szmcdull/huobi_golang/pkg/model/wallet"
)

// Responsible to operate wallet
type SubUserClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

// Initializer
func (p *SubUserClient) Init(accessKey string, secretKey string, host string) *SubUserClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

// Parent user query sub user deposit address of corresponding chain, for a specific crypto currency (except IOTA)
func (p *SubUserClient) CreateSubUser(request subuser.CreateSubUserRequest) ([]subuser.UserData, error) {
	postBody, jsonErr := model.ToJson(request)

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/creation", nil)
	postResp, postErr := internal.HttpPost(url, string(postBody))
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.CreateSubUserResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(postResp)
}

// CreateAPIKey Parent user query to create the API key of the sub user
func (p *SubUserClient) CreateAPIKey(request subuser.CreateAPIKeyRequest) (*subuser.APIKeyData, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/api-key-generation", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.CreateAPIKeyResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code != 200 {
		return nil, errors.New(postResp)
	}
	return result.Data, nil
}

// ModifyAPIKey Parent user query to modify the API key of the sub user
func (p *SubUserClient) ModifyAPIKey(request subuser.ModifyAPIKeyRequest) (*subuser.APIKeyData, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/api-key-modification", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.CreateAPIKeyResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code != 200 {
		return nil, errors.New(postResp)
	}
	return result.Data, nil
}

// DeleteAPIKey Parent user query to delete the API key of the sub user
func (p *SubUserClient) DeleteAPIKey(request subuser.DeleteAPIKeyRequest) error {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/api-key-deletion", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return postErr
	}

	result := subuser.DeleteAPIKeyResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return jsonErr
	}
	if result.Code != 200 {
		return errors.New(postResp)
	}
	return nil
}

// Lock or unlock a specific user
func (p *SubUserClient) SubUserManagement(request subuser.SubUserManagementRequest) (*subuser.SubUserManagement, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/management", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.SubUserManagementResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code != 200 {
		return nil, errors.New(postResp)
	}
	return result.Data, nil
}

// Set Tradable Market for Sub Users
func (p *SubUserClient) SetSubUserTradableMarket(request subuser.SetSubUserTradableMarketRequest) ([]subuser.TradableMarket, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/tradable-market", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.SetSubUserTradableMarketResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code != 200 {
		return nil, errors.New(postResp)
	}

	return result.Data, nil
}

// Set Asset Transfer Permission for Sub Users
func (p *SubUserClient) SetSubUserTransferability(request subuser.SetSubUserTransferabilityRequest) ([]subuser.Transferability, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return nil, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v2/sub-user/transferability", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := subuser.SetSubUserTransferabilityResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code != 200 {
		return nil, errors.New(postResp)
	}

	return result.Data, nil
}

// Transfer asset between parent and sub account
func (p *SubUserClient) SubUserTransfer(request subuser.SubUserTransferRequest) (int64, error) {
	postBody, jsonErr := model.ToJson(request)
	if jsonErr != nil {
		return 0, jsonErr
	}

	url := p.privateUrlBuilder.Build("POST", "/v1/subuser/transfer", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return 0, postErr
	}

	result := new(subuser.SubUserTransferResponse)
	if jsonErr = json.Unmarshal([]byte(postResp), result); jsonErr != nil {
		return 0, jsonErr
	}

	if result.Status == "ok" && result.Data != nil {
		return *result.Data, nil
	}

	return 0, errors.New(postResp)
}

// Parent user query sub user deposit address of corresponding chain, for a specific crypto currency (except IOTA)
func (p *SubUserClient) GetSubUserDepositAddress(subUid int64, currency string) ([]wallet.DepositAddress, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("subUid", strconv.FormatInt(subUid, 10))
	request.AddParam("currency", currency)

	url := p.privateUrlBuilder.Build("GET", "/v2/sub-user/deposit-address", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := wallet.GetDepositAddressResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Parent user query sub user deposits history
func (p *SubUserClient) QuerySubUserDepositHistory(subUid int64, optionalRequest subuser.QuerySubUserDepositHistoryOptionalRequest) ([]subuser.DepositHistory, error) {
	request := new(model.GetRequest).Init()

	request.AddParam("subUid", strconv.FormatInt(subUid, 10))

	if optionalRequest.Currency != "" {
		request.AddParam("currency", optionalRequest.Currency)
	}
	if optionalRequest.StartTime != 0 {
		request.AddParam("startTime", strconv.FormatInt(optionalRequest.StartTime, 10))
	}
	if optionalRequest.EndTime != 0 {
		request.AddParam("endTime", strconv.FormatInt(optionalRequest.EndTime, 10))
	}
	if optionalRequest.Sort != "" {
		request.AddParam("sort", optionalRequest.Sort)
	}
	if optionalRequest.Limit != "" {
		request.AddParam("limit", optionalRequest.Limit)
	}
	if optionalRequest.FromId != 0 {
		request.AddParam("fromId", strconv.FormatInt(optionalRequest.FromId, 10))
	}

	url := p.privateUrlBuilder.Build("GET", "/v2/sub-user/query-deposit", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := subuser.QuerySubUserDepositHistoryResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}

// Returns the aggregated balance from all the sub-users
func (p *SubUserClient) GetSubUserAggregateBalance() ([]account.AccountBalance, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/subuser/aggregate-balance", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetSubUserAggregateBalanceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

// Returns the balance of a sub-account specified by sub-uid
func (p *SubUserClient) GetSubUserAccount(subUid int64) ([]account.SubUserAccount, error) {
	url := p.privateUrlBuilder.Build("GET", fmt.Sprintf("/v1/account/accounts/%d", subUid), nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}
	result := account.GetSubUserAccountResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {
		return result.Data, nil
	}

	return nil, errors.New(getResp)
}

func (p *SubUserClient) GetUid() (int64, error) {
	url := p.privateUrlBuilder.Build("GET", "/v2/user/uid", nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return 0, getErr
	}
	result := account.GetUidResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return 0, jsonErr
	}
	if result.Code == 200 && result.Data != 0 {
		return result.Data, nil
	}
	return 0, errors.New(getResp)
}

func (p *SubUserClient) GetSubUsersAccountList(subUid int64) (*subuser.AccountData, error) {
	request := new(model.GetRequest).Init()
	request.AddParam("subUid", strconv.FormatInt(subUid, 10))
	url := p.privateUrlBuilder.Build("GET", "/v2/sub-user/account-list", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := subuser.GetSubUsersAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Code == 200 && result.Data != nil {
		return result.Data, nil
	}
	return nil, errors.New(getResp)
}
