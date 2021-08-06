package subuser

type CreateAPIKeyRequest struct {
	OtpToken string `json:"otpToken"`
	SubUid int64 `json:"sub_uid"`
	Note string `json:"note"`
	IpAddresses *string `json:"ipAddresses"`
}
