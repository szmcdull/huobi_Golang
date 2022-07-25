package linearswap

type AccountFinancialRecordsResponse struct {
	Status string                       `json:"status"`
	Data   *AccountFinancialRecordsData `json:"data"`
}

type AccountFinancialRecordsData struct {
	TotalPage        int                       `json:"total_page"`
	CurrentPage      int                       `json:"current_page"`
	TotalSize        int                       `json:"total_size"`
	FinancialRecords []*AccountFinancialRecord `json:"financial_record"`
}
