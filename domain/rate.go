package domain

type (
	// Rate contains exchange rate for some fiat currency
	Rate struct {
		M      float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	}

	// Rates is a map of Rate where key is a fiat currency name
	Rates map[FiatCurrency]*Rate

	// ExchangeRateConverter converts wallet's balance to fiat currency
	ExchangeRateConverter interface {
		Convert(*Wallet, FiatCurrency) (float64, error)
	}
)
