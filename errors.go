package sjwt

import "errors"

var (
	// ErrTokenInvalid is an error string clarifying
	// the provided token is an invalid format
	ErrTokenInvalid = errors.New("Token is invalid")

	// ErrTokenHasExpired is an error string clarifying
	// the current unix timestamp has exceed the exp unix timestamp
	ErrTokenHasExpired = errors.New("Token has expired")

	// ErrTokenNotYetValid is an error string clarifying
	// the current unix timestamp has not exceeded the nbf unix timestamp
	ErrTokenNotYetValid = errors.New("Token is not yet valid")
)
