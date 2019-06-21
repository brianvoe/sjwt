package sjwt

const (
	// TokenID is a unique identifier for this token
	TokenID = "jti"

	// Issuer is the principal that issued the token
	Issuer = "iss"

	// Audience identifies the recipents the token is intended for
	Audience = "aud"

	// Subject is the subject of the token
	Subject = "sub"

	// IssuedAt is a timesatamp for when the token was issued
	IssuedAt = "iat"

	// ExpiresAt is a timestamp for when the token should expire
	ExpiresAt = "exp"

	// NotBefore is a timestamp for which this token should not be excepted until
	NotBefore = "nbf"
)

// SetTokenID will set a random uuid v4 id
func (c Claims) SetTokenID() {
	c[TokenID] = UUID()
}

// GetTokenID will get the id set on the Claims
func (c Claims) GetTokenID() string {
	if val, ok := c[TokenID]; ok {
		return val.(string)
	}
	return ""
}

// SetIssuer will set a string value for the issuer
func (c Claims) SetIssuer(issuer string) {
	c[Issuer] = issuer
}

// GetIssuer will get the issuer set on the Claims
func (c Claims) GetIssuer() string {
	if val, ok := c[Issuer]; ok {
		return val.(string)
	}
	return ""
}

// SetAudience will set a string value for the audience
func (c Claims) SetAudience(audience string) {
	c[Audience] = audience
}

// GetAudience will get the audience set on the Claims
func (c Claims) GetAudience() string {
	if val, ok := c[Audience]; ok {
		return val.(string)
	}
	return ""
}

// SetSubject will set a subject value
func (c Claims) SetSubject(subject string) {
	c[Subject] = subject
}

// GetSubject will get the subject set on the Claims
func (c Claims) GetSubject() string {
	if val, ok := c[Subject]; ok {
		return val.(string)
	}
	return ""
}

// SetIssuedAt will set an issued at timestamp in nanoseconds
func (c Claims) SetIssuedAt(issuedAt int) {
	c[IssuedAt] = issuedAt
}

// GetIssuedAt will get the issued at timestamp set on the Claims
func (c Claims) GetIssuedAt() int {
	if val, ok := c[IssuedAt]; ok {
		return val.(int)
	}
	return 0
}

// SetExpiresAt will set an expires at timestamp in nanoseconds
func (c Claims) SetExpiresAt(expiresAt int) {
	c[ExpiresAt] = expiresAt
}

// GetExpiresAt will get the expires at timestamp set on the Claims
func (c Claims) GetExpiresAt() int {
	if val, ok := c[ExpiresAt]; ok {
		return val.(int)
	}
	return 0
}

// SetNotBeforeAt will set an not before at timestamp in nanoseconds
func (c Claims) SetNotBeforeAt(notbeforeAt int) {
	c[NotBefore] = notbeforeAt
}

// GetNotBeforeAt will get the not before at timestamp set on the Claims
func (c Claims) GetNotBeforeAt() int {
	if val, ok := c[NotBefore]; ok {
		return val.(int)
	}
	return 0
}
