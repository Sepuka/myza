package errors

import "errors"

var (
	UnknownCurrency = errors.New(`unknkown currency`)
)

// NewUnknownCurrencyError instance an InvalidJson error
func NewUnknownCurrencyError(msg string, originalErr error) MyzaError {
	return MyzaError{
		err:           UnknownCurrency,
		message:       msg,
		originalError: originalErr,
	}
}
