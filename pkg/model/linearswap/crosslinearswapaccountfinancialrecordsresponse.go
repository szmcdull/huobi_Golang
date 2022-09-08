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

type CrossLinearSwapAccountFinancialRecordsViaMultipleFieldsResponse struct {
	Status string                                                       `json:"status"`
	Ts     int64                                                        `json:"ts"`
	Data   *CrossLinearSwapAccountFinancialRecordsViaMultipleFieldsData `json:"data"`
}

type CrossLinearSwapAccountFinancialRecordsViaMultipleFieldsData struct {
	FinancialRecords []*CrossLinearSwapAccountFinancialRecord `json:"financial_record"`
	RemainSize       int                                      `json:"remain_size"`
	NextId           interface{}                              `json:"next_id"`
}

type CrossLinearSwapAccountFinancialRecordsV3Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Ts   int64  `json:"ts"`
	Data []*CrossLinearSwapAccountFinancialRecordV3
}
