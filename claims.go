package sjwt

import (
	"encoding/json"
	"time"
)

// Claims is the main container for our body information
type Claims map[string]interface{}

// New will initiate a new claims
func New() *Claims {
	return &Claims{}
}

// ToClaims takes in an interface and unmarshals it to claims
func ToClaims(struc interface{}) (Claims, error) {
	strucBytes, _ := json.Marshal(struc)
	var claims Claims
	err := json.Unmarshal(strucBytes, &claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// ToStruct takes your claims and sets value to struct
func (c Claims) ToStruct(struc interface{}) error {
	claimsBytes, _ := json.Marshal(c)
	err := json.Unmarshal(claimsBytes, struc)
	if err != nil {
		return err
	}

	return nil
}

// Validate checks expiration and not before times
func (c Claims) Validate() error {
	now := time.Now().Unix()

	// Check if not before at is set and if current time hasnt started yet
	if c.Has(NotBeforeAt) {
		nbf, _ := c.GetNotBeforeAt()
		if now < nbf {
			return ErrTokenNotYetValid
		}
	}

	// Check if expiration at is set and if current time is passed
	if c.Has(ExpiresAt) {
		exp, _ := c.GetExpiresAt()
		if now >= exp {
			return ErrTokenHasExpired
		}
	}

	return nil
}
