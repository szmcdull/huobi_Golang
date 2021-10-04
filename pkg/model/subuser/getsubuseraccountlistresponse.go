package subuser

type GetSubUsersAccountListResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *AccountData `json:"data,omitempty"`
}

type AccountData struct {
	Uid        int64      `json:"uid"`
	DeductMode string     `json:"deductMode"`
	List       []*Account `json:"list,omitempty"`
}

type AccountId struct {
	AccountId     int64  `json:"accountId"`
	AccountStatus string `json:"accountStatus"`
}

type Account struct {
	AccountType   string       `json:"accountType"`
	Activation    string       `json:"activation"`
	Transferrable bool         `json:"transferrable,omitempty"`
	AccountIds    []*AccountId `json:"accountIds,omitempty"`
}
