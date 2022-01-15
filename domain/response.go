package domain

import "encoding/json"

const (
	MinorDiv = 100000000
)

type Wallet struct {
	Hash160       string  `json:"hash160"`
	Address       string  `json:"address"`
	NTx           int     `json:"n_tx"`
	NUnredeemed   int     `json:"n_unredeemed"`
	TotalReceived int     `json:"total_received"`
	TotalSent     int     `json:"total_sent"`
	FinalBalance  float64 `json:"final_balance"`
}

func NewWallet() *Wallet {
	return &Wallet{}
}

func (r *Wallet) BalanceToBTC() float64 {
	return r.FinalBalance / MinorDiv
}

func (r *Wallet) MarshalBinary() (data []byte, err error) {
	return json.Marshal(r)
}
