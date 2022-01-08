package domain

type AddrResponse struct {
	Hash160       string `json:"hash160"`
	Address       string `json:"address"`
	NTx           int    `json:"n_tx"`
	NUnredeemed   int    `json:"n_unredeemed"`
	TotalReceived int    `json:"total_received"`
	TotalSent     int    `json:"total_sent"`
	FinalBalance  int    `json:"final_balance"`
}

func NewAddrResponse() *AddrResponse {
	return &AddrResponse{}
}
