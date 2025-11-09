package sjwt

import "errors"

var (
	// ErrNotFound is an error string clarifying
	// that the attempted key does not exist in the claims
	ErrNotFound = errors.New("claim key not found in claims")

	// ErrClaimValueInvalid is an error string clarifying
	// that the attempt to retrieve a value could not be properly converted
	ErrClaimValueInvalid = errors.New("claim value invalid")

	// ErrTokenInvalid is an error string clarifying
	// the provided token is an invalid format
	ErrTokenInvalid = errors.New("token is invalid")

	// ErrTokenHasExpired is an error string clarifying
	// the current unix timestamp has exceed the exp unix timestamp
	ErrTokenHasExpired = errors.New("token has expired")

	// ErrTokenNotYetValid is an error string clarifying
	// the current unix timestamp has not exceeded the nbf unix timestamp
	ErrTokenNotYetValid = errors.New("token is not yet valid")

	// ErrTokenSignatureInvalid clarifies the token signature did not match the expected value
	ErrTokenSignatureInvalid = errors.New("token signature invalid")

	// ErrTokenHeaderInvalid clarifies the token header is malformed or unsupported
	ErrTokenHeaderInvalid = errors.New("token header invalid")

	// ErrTokenAlgorithmMismatch clarifies that the token algorithm does not match the supported algorithm
	ErrTokenAlgorithmMismatch = errors.New("token algorithm mismatch")

	// ErrSecretTooShort clarifies that the provided secret is weaker than the minimum required length
	ErrSecretTooShort = errors.New("secret key too short; use at least 32 random bytes")
)
