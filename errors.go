package sjwt

import "errors"

var (
	// ErrTokenHasExpired is an error string clarifying
	// the current unix timestamp has exceed the exp unix timestamp
	ErrTokenHasExpired = errors.New("Token has expired")

	// ErrTokenNotYetValid is an error string clarifying
	// the current unix timestamp has not exceeded the nbf unix timestamp
	ErrTokenNotYetValid = errors.New("Token is not yet valid")
)
