package btc

import "github.com/sepuka/myza/domain"

func NewAddressGeneratorContext(
	currency domain.CryptoCurrency,
	userId uint32,
) *domain.AddressGeneratorContext {
	return &domain.AddressGeneratorContext{
		Currency: currency,
		UserId:   userId,
	}
}
