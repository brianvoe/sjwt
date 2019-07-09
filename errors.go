package sjwt

import "errors"

var (
	// ErrNotFound is an error string clarifying
	// that the attempted key does not exist in the claims
	ErrNotFound = errors.New("Claim key not found in claims")

	// ErrClaimValueInvalid is an error string clarifying
	// that the attempt to retrieve a value could not be properly converted
	ErrClaimValueInvalid = errors.New("Claim value invalid")

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
