package sjwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

const (
	jwtType             = "JWT"
	jwtAlgorithm        = "HS256"
	minSecretLength     = 32
	tokenSegments       = 3
	headerSegmentIdx    = 0
	payloadSegmentIdx   = 1
	signatureSegmentIdx = 2
)

type jwtHeader struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

// Generate takes in claims and a secret and outputs jwt token
func (c Claims) Generate(secret []byte) (string, error) {
	if len(secret) < minSecretLength {
		return "", ErrSecretTooShort
	}
	// Encode header and claims
	headerEnc, err := json.Marshal(jwtHeader{Typ: jwtType, Alg: jwtAlgorithm})
	if err != nil {
		return "", err
	}

	claimsEnc, err := json.Marshal(c)
	if err != nil {
		return "", err
	}

	headerEncoded := make([]byte, base64.RawURLEncoding.EncodedLen(len(headerEnc)))
	base64.RawURLEncoding.Encode(headerEncoded, headerEnc)

	payloadEncoded := make([]byte, base64.RawURLEncoding.EncodedLen(len(claimsEnc)))
	base64.RawURLEncoding.Encode(payloadEncoded, claimsEnc)

	unsignedLen := len(headerEncoded) + 1 + len(payloadEncoded)
	unsigned := make([]byte, unsignedLen)
	copy(unsigned, headerEncoded)
	unsigned[len(headerEncoded)] = '.'
	copy(unsigned[len(headerEncoded)+1:], payloadEncoded)

	// Sign with sha 256
	mac := hmac.New(sha256.New, secret)
	mac.Write(unsigned)

	sig := mac.Sum(nil)
	signatureEncoded := make([]byte, base64.RawURLEncoding.EncodedLen(len(sig)))
	base64.RawURLEncoding.Encode(signatureEncoded, sig)

	token := make([]byte, unsignedLen+1+len(signatureEncoded))
	copy(token, unsigned)
	token[unsignedLen] = '.'
	copy(token[unsignedLen+1:], signatureEncoded)

	return string(token), nil
}

// Parse takes in the token string and returns the claims payload (without verifying the signature)
func Parse(tokenStr string) (Claims, error) {
	tokenArray := splitToken(tokenStr)
	if len(tokenArray) != tokenSegments {
		return nil, ErrTokenInvalid
	}

	if err := validateHeader(tokenArray[headerSegmentIdx]); err != nil {
		return nil, err
	}

	payload := tokenArray[payloadSegmentIdx]

	decodedLen := base64.RawURLEncoding.DecodedLen(len(payload))
	claimsByte := make([]byte, decodedLen)
	n, err := base64.RawURLEncoding.Decode(claimsByte, []byte(payload))
	if err != nil {
		return nil, err
	}
	claimsByte = claimsByte[:n]

	var claims Claims
	err = json.Unmarshal(claimsByte, &claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// Verify will take in the token string and secret and identify the signature matches
func Verify(tokenStr string, secret []byte) bool {
	token := splitToken(tokenStr)
	if len(token) != tokenSegments {
		return false
	}
	if err := verifySignature(token, secret); err != nil {
		return false
	}
	return true
}

func verifySignature(token []string, secret []byte) error {
	if len(secret) < minSecretLength {
		return ErrSecretTooShort
	}

	header := token[headerSegmentIdx]
	payload := token[payloadSegmentIdx]

	unsignedLen := len(header) + 1 + len(payload)
	unsigned := make([]byte, unsignedLen)
	copy(unsigned, header)
	unsigned[len(header)] = '.'
	copy(unsigned[len(header)+1:], payload)

	mac := hmac.New(sha256.New, secret)
	mac.Write(unsigned)
	sig, err := base64.RawURLEncoding.DecodeString(token[signatureSegmentIdx])
	if err != nil {
		return ErrTokenSignatureInvalid
	}
	if !hmac.Equal(sig, mac.Sum(nil)) {
		return ErrTokenSignatureInvalid
	}

	return nil
}

func validateHeader(segment string) error {
	headerBytes, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil {
		return ErrTokenHeaderInvalid
	}

	var header jwtHeader
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		return ErrTokenHeaderInvalid
	}

	if header.Typ != "" && header.Typ != jwtType {
		return ErrTokenHeaderInvalid
	}
	if header.Alg != jwtAlgorithm {
		return ErrTokenAlgorithmMismatch
	}
	return nil
}

func splitToken(token string) []string {
	if token == "" {
		return []string{}
	}
	parts := make([]string, 0, tokenSegments)
	start := 0
	for i := 0; i < len(token); i++ {
		if token[i] == '.' {
			parts = append(parts, token[start:i])
			start = i + 1
		}
	}
	parts = append(parts, token[start:])
	return parts
}
