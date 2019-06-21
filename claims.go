package sjwt

import "encoding/json"

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
