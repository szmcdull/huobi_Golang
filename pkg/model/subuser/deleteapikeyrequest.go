package subuser

type DeleteAPIKeyRequest struct {
	SubUid int64 `json:"subUid"`
	AccessKey string `json:"accessKey"`
}
