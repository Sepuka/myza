package domain

type (
	// CryptoCurrency is a cryptocurrency type
	CryptoCurrency string
	// FiatCurrency is a fiat currency type
	FiatCurrency string
)

const (
	Rub FiatCurrency   = `RUB`
	Btc CryptoCurrency = `btc`
)
