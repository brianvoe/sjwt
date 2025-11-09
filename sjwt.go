package sjwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

	jwtStr := base64.RawURLEncoding.EncodeToString(headerEnc) + "." + base64.RawURLEncoding.EncodeToString(claimsEnc)

	// Sign with sha 256
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(jwtStr))

	return fmt.Sprintf("%s.%s", jwtStr, base64.RawURLEncoding.EncodeToString(mac.Sum(nil))), nil
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

	claimsByte, err := base64.RawURLEncoding.DecodeString(tokenArray[payloadSegmentIdx])
	if err != nil {
		return nil, err
	}

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

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(token[headerSegmentIdx] + "." + token[payloadSegmentIdx]))
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
