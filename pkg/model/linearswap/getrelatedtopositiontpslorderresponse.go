package linearswap

type GetRelatedToPositionTPSLOrderResponse struct {
	Status string        `json:"status"`
	Data   *TriggerOrder `json:"data"`
}
