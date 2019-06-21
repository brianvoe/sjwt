package sjwt

// AddClaim adds a name/value to claims
func (c Claims) AddClaim(name string, value interface{}) { c[name] = value }

// HasClaim will let you know whether or not a claim exists
func (c Claims) HasClaim(name string) bool { _, ok := c[name]; return ok }

// GetClaim gets claim value
func (c Claims) GetClaim(name string) interface{} { return c[name] }

// GetClaimStr will get the string value on the Claims
func (c Claims) GetClaimStr(name string) string {
	if val, ok := c[name]; ok {
		return val.(string)
	}
	return ""
}

// GetClaimInt will get the int value on the Claims
func (c Claims) GetClaimInt(name string) int {
	if val, ok := c[name]; ok {
		return val.(int)
	}
	return 0
}

// GetClaimFloat will get the float value on the Claims
func (c Claims) GetClaimFloat(name string) float32 {
	if val, ok := c[name]; ok {
		return val.(float32)
	}
	return 0
}
