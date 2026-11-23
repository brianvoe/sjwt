package sjwt

import (
	"strconv"
	"testing"
	"time"
)

type testStruc struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func TestToClaims(t *testing.T) {
	test := testStruc{
		FirstName: "Billy",
		LastName:  "Mister",
	}

	claims, err := ToClaims(test)
	if err != nil {
		t.Error("Error ToClaims: ", err)
	}
	if !claims.Has("first_name") {
		t.Error("Tried to get claim from struct. Non found")
	}

}

func TestToStruct(t *testing.T) {
	claims := New()
	claims.Set("first_name", "Billy")
	claims.Set("last_name", "Mister")

	// Try to set claims into struct
	var test testStruc
	if err := claims.ToStruct(&test); err != nil {
		t.Fatalf("ToStruct returned error: %v", err)
	}

	if test.FirstName != "Billy" {
		t.Error("Tried to get first name from test struct after running ToStruct and it failed")
	}
}

func TestValidate(t *testing.T) {
	// Validate just the claim
	claims := New()
	claims.SetIssuedAt(time.Now())
	claims.SetNotBeforeAt(time.Now())
	claims.SetExpiresAt(time.Now().Add(time.Hour))
	err := claims.Validate()
	if err != nil {
		t.Error("Validate was not successful when it should be")
	}

	// Validate on parsed claims
	token, err := claims.Generate(secretKey)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	parsedClaims, err := Parse(token)
	if err != nil {
		t.Error("Error parsing token: ", err)
	}

	err = parsedClaims.Validate()
	if err != nil {
		t.Error("Validate was not successful on parsed claims when it should be")
	}
}

func TestValidateExp(t *testing.T) {
	// Succes
	claims := New()
	claims.SetExpiresAt(time.Now().Add(time.Hour))
	err := claims.Validate()
	if err != nil {
		t.Error("Validate was not successful when it should be")
	}

	// Error
	claims.SetExpiresAt(time.Now().Add(time.Hour * -1))
	err = claims.Validate()
	if err != ErrTokenHasExpired {
		t.Error("Token should have expired")
	}
}

func TestValidateNotBefore(t *testing.T) {
	// Succes
	claims := New()
	claims.SetNotBeforeAt(time.Now())
	err := claims.Validate()
	if err != nil {
		t.Error("Validate was not successful when it should be")
	}

	// Error
	claims.SetNotBeforeAt(time.Now().Add(time.Hour))
	err = claims.Validate()
	if err != ErrTokenNotYetValid {
		t.Error("Token should have failed due to token not being valid yet")
	}
}

// A not before or an expiry that cannot be read as a time is an error: the token must not pass because of it
func TestValidateMalformedTimes(t *testing.T) {
	values := map[string]any{
		"text":   "soon",
		"bool":   true,
		"null":   nil,
		"list":   []any{1},
		"object": map[string]any{},
	}

	for _, name := range []string{NotBeforeAt, ExpiresAt} {
		for label, value := range values {
			claims := New()
			claims.Set(name, value)

			token, err := claims.Generate(secretKey)
			if err != nil {
				t.Fatalf("Generate returned error: %v", err)
			}

			parsed, err := Parse(token)
			if err != nil {
				t.Fatalf("Parse returned error: %v", err)
			}

			if err := claims.Validate(); err != ErrClaimValueInvalid {
				t.Errorf("%s as %s: got %v, want %v", name, label, err, ErrClaimValueInvalid)
			}

			if err := parsed.Validate(); err != ErrClaimValueInvalid {
				t.Errorf("%s as %s after Parse: got %v, want %v", name, label, err, ErrClaimValueInvalid)
			}
		}
	}
}

// A time can still be a number, a number with a fraction or the digits in a string
func TestValidateTimeFormats(t *testing.T) {
	claims := New()
	claims.Set(NotBeforeAt, strconv.FormatInt(time.Now().Add(-time.Hour).Unix(), 10))
	claims.Set(ExpiresAt, float64(time.Now().Add(time.Hour).Unix())+0.5)
	if err := claims.Validate(); err != nil {
		t.Errorf("Validate returned %v for times that are readable", err)
	}
}
