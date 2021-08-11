package subuser

type CreateAPIKeyRequest struct {
	OtpToken    string  `json:"otpToken"`
	SubUid      int64   `json:"subUid"`
	Note        string  `json:"note"`
	Permission  string  `json:"permission"`
	IpAddresses *string `json:"ipAddresses"`
}
