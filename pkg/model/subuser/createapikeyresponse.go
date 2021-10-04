package subuser

type CreateAPIKeyResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    *APIKeyData `json:"data"`
}

type APIKeyData struct {
	Note        string `json:"note"`
	AccessKey   string `json:"accessKey"`
	SecretKey   string `json:"secretKey"`
	Permission  string `json:"permission"`
	IpAddresses string `json:"ip_addresses"`
}
