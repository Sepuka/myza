package errors

import "errors"

var (
	UnknownCurrency        = errors.New(`unknown currency`)
	InvalidNet             = errors.New(`invalid blockchain net`)
	OauthTokenError        = errors.New(`oauth token error`)
	BlockchainBalanceError = errors.New(`get balance error`)
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

func NewOauthTokenError(msg string, err error) MyzaError {
	return MyzaError{
		err:           OauthTokenError,
		message:       msg,
		originalError: err,
	}
}

func NewBlockchainBalanceError(msg string) MyzaError {
	return MyzaError{
		err:     BlockchainBalanceError,
		message: msg,
	}
}
