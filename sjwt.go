package sjwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// New takes in claims and a secret
func New(claims interface{}, secret []byte) (string, error) {
	claimsMarshaled, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	encodeClaims := base64.RawURLEncoding.EncodeToString(claimsMarshaled)

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(encodeClaims))

	return fmt.Sprintf("%s.%s", encodeClaims, base64.RawURLEncoding.EncodeToString(mac.Sum(nil))), nil
}

// ParseClaims will take in the token string grab the body and unmarshal into claims interface
func ParseClaims(tokenStr string, claims interface{}) error {
	claimsStr := strings.Split(tokenStr, ".")[0]

	claimsByte, err := base64.RawURLEncoding.DecodeString(claimsStr)
	if err != nil {
		return err
	}

	err = json.Unmarshal(claimsByte, claims)
	if err != nil {
		return err
	}

	return nil
}

// Verify will take in the token string and secret and identify the signature matches
func Verify(tokenStr string, secret []byte) bool {
	token := strings.Split(tokenStr, ".")
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(token[0]))
	return hmac.Equal([]byte(token[1]), mac.Sum(nil))

}
