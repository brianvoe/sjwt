package sjwt

import (
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
	claims.ToStruct(&test)

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
	token := claims.Generate([]byte(secretKey))
	parsedClaims, err := Parse(token)
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
