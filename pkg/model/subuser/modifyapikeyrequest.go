package subuser

type ModifyAPIKeyRequest struct {
	SubUid int64 `json:"subUid"`
	AccessKey string `json:"accessKey"`
	Note *string `json:"note"`
	Permission *string `json:"permission"`
	IpAddresses *string `json:"ipAddresses"`
}


