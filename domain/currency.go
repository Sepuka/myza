package domain

type (
	// CryptoCurrency is a crypto currency type
	CryptoCurrency string
	// FiatCurrency is a fiat currency type
	FiatCurrency string
)

const (
	Rub FiatCurrency   = `RUB`
	Btc CryptoCurrency = `btc`
)
