package sjwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// Generate takes in claims and a secret and outputs jwt token
func (c Claims) Generate(secret []byte) string {
	// Encode header and claims
	headerEnc, _ := json.Marshal(map[string]string{"typ": "JWT", "alg": "HS256"})
	claimsEnc, _ := json.Marshal(c)
	jwtStr := fmt.Sprintf(
		"%s.%s",
		base64.RawURLEncoding.EncodeToString(headerEnc),
		base64.RawURLEncoding.EncodeToString(claimsEnc),
	)

	// Sign with sha 256
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(jwtStr))

	return fmt.Sprintf("%s.%s", jwtStr, base64.RawURLEncoding.EncodeToString(mac.Sum(nil)))
}

// Parse will take in the token string grab the body and unmarshal into claims interface
func Parse(tokenStr string) (Claims, error) {
	tokenArray := strings.Split(tokenStr, ".")

	// Make sure token array contains 3 parts
	if len(tokenArray) != 3 {
		return nil, ErrTokenInvalid
	}

	claimsByte, err := base64.RawURLEncoding.DecodeString(tokenArray[1])
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
	token := strings.Split(tokenStr, ".")
	if len(token) != 3 {
		return false
	}
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(fmt.Sprintf("%s.%s", token[0], token[1])))
	sig, _ := base64.RawURLEncoding.DecodeString(token[2])
	return hmac.Equal(sig, mac.Sum(nil))
}
