package linearswap

type CrossLinearSwapAccountFinancialRecordsResponse struct {
	Status string                                      `json:"status"`
	Data   *CrossLinearSwapAccountFinancialRecordsData `json:"data"`
}

type CrossLinearSwapAccountFinancialRecordsData struct {
	TotalPage        int                                      `json:"total_page"`
	CurrentPage      int                                      `json:"current_page"`
	TotalSize        int                                      `json:"total_size"`
	FinancialRecords []*CrossLinearSwapAccountFinancialRecord `json:"financial_record"`
}
