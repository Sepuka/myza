package domain

const (
	MinorDiv = 100000000
)

type AddrResponse struct {
	Hash160       string  `json:"hash160"`
	Address       string  `json:"address"`
	NTx           int     `json:"n_tx"`
	NUnredeemed   int     `json:"n_unredeemed"`
	TotalReceived int     `json:"total_received"`
	TotalSent     int     `json:"total_sent"`
	FinalBalance  float64 `json:"final_balance"`
}

func NewAddrResponse() *AddrResponse {
	return &AddrResponse{}
}

func (r *AddrResponse) BalanceToBTC() float64 {
	return r.FinalBalance / MinorDiv
}
