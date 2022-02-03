package errors

import "errors"

var (
	UnknownCurrency = errors.New(`unknown currency`)
	InvalidNet      = errors.New(`invalid blockchain net`)
)

// NewUnknownCurrencyError instance an InvalidJson error
func NewUnknownCurrencyError(msg string, originalErr error) MyzaError {
	return MyzaError{
		err:           UnknownCurrency,
		message:       msg,
		originalError: originalErr,
	}
}

// NewInvalidBlockchainNetError instance an InvalidJson error
func NewInvalidBlockchainNetError(msg string) MyzaError {
	return MyzaError{
		err:     InvalidNet,
		message: msg,
	}
}
