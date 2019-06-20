package sjwt

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Token struct {
	Header Header
	Claims interface{}
}

type Header struct {
	Type      string `json:"typ"`
	Algorithm string `json:"alg"`
}

// New takes in claims and a secret
func New(claims interface{}, secret []byte) (string, error) {
	// Marshal Header
	headerBytes, err := json.Marshal(Header{
		Type:      "JWT",
		Algorithm: "HS512",
	})
	if err != nil {
		return "", err
	}

	// Marshal Claims
	claimsBytes, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	// Put together jwt string
	jwtStr := fmt.Sprintf(
		"%s.%s",
		base64.RawURLEncoding.EncodeToString(headerBytes),
		base64.RawURLEncoding.EncodeToString(claimsBytes),
	)

	// Create New hmac and write jwt string to it
	mac := hmac.New(sha512.New, secret)
	mac.Write([]byte(jwtStr))

	return fmt.Sprintf("%s.%s", jwtStr, base64.RawURLEncoding.EncodeToString(mac.Sum(nil))), nil

}

// Parse will take in the token string grab the body and unmarshal into claims interface
func Parse(tokenStr string, claims interface{}) (*Token, error) {
	tokenArray := strings.Split(tokenStr, ".")

	token := Token{}

	headerByte, err := base64.RawURLEncoding.DecodeString(tokenArray[0])
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(headerByte, token.Header)
	if err != nil {
		return nil, err
	}

	claimsByte, err := base64.RawURLEncoding.DecodeString(tokenArray[1])
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(claimsByte, token.Claims)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// Verify will take in the token string and secret and identify the signature matches
func (t *Token) Verify(tokenStr string, secret []byte) bool {
	token := strings.Split(tokenStr, ".")
	mac := hmac.New(sha512.New, secret)
	mac.Write([]byte(token[0]))
	return hmac.Equal([]byte(token[1]), mac.Sum(nil))

}
