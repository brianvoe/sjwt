package sjwt

import (
	"slices"
	"testing"
	"time"
)

func TestTokenId(t *testing.T) {
	claims := New()
	claims.SetTokenID()
	tokenID, _ := claims.GetTokenID()
	if tokenID == "" {
		t.Error("token id was not set")
	}
	claims.DeleteTokenID()
	if claims.Has(TokenID) {
		t.Error("token id should have been deleted")
	}
	tokenID, _ = claims.GetTokenID()
	if tokenID != "" {
		t.Error("should have gotten blank value")
	}
}

func TestIssuer(t *testing.T) {
	claims := New()
	claims.SetIssuer("Google")
	issuer, _ := claims.GetIssuer()
	if issuer != "Google" {
		t.Error("issuer was not set")
	}
	claims.DeleteIssuer()
	if claims.Has(Issuer) {
		t.Error("issuer should have been deleted")
	}
	issuer, _ = claims.GetIssuer()
	if issuer != "" {
		t.Error("should have gotten blank value")
	}
}

func TestAudience(t *testing.T) {
	claims := New()
	claims.SetAudience([]string{"Google", "Facebook"})
	audience, _ := claims.GetAudience()
	if len(audience) != 2 || audience[0] != "Google" || audience[1] != "Facebook" {
		t.Error("audience was not set")
	}
	claims.DeleteAudience()
	if claims.Has(Audience) {
		t.Error("audience should have been deleted")
	}
	audience, _ = claims.GetAudience()
	if len(audience) != 0 {
		t.Error("should have gotten empty string array")
	}
}

// The audience of a token that went through Parse is a list of values of any type, not a []string
func TestAudienceAfterParse(t *testing.T) {
	for _, want := range [][]string{{"Google"}, {"Google", "Facebook"}, {}, {"Zoë", "Ünal", "日本"}} {
		claims := New()
		claims.SetAudience(want)
		token, err := claims.Generate(secretKey)
		if err != nil {
			t.Fatalf("Generate returned error: %v", err)
		}

		parsed, err := Parse(token)
		if err != nil {
			t.Fatalf("Parse returned error: %v", err)
		}

		audience, err := parsed.GetAudience()
		if err != nil || !slices.Equal(audience, want) {
			t.Errorf("audience %v after Parse: got %v, %v", want, audience, err)
		}
	}
}

// Parse gives back every registered claim that Generate was given
func TestRegisteredClaimsAfterParse(t *testing.T) {
	issuedAt := time.Now().Add(-time.Minute)
	notBeforeAt := time.Now().Add(-30 * time.Second)
	expiresAt := time.Now().Add(time.Hour)

	claims := New()
	claims.SetTokenID()
	claims.SetIssuer("Google")
	claims.SetAudience([]string{"Google", "Facebook"})
	claims.SetSubject("Billy")
	claims.SetIssuedAt(issuedAt)
	claims.SetNotBeforeAt(notBeforeAt)
	claims.SetExpiresAt(expiresAt)
	tokenID, _ := claims.GetTokenID()

	token, err := claims.Generate(secretKey)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	parsed, err := Parse(token)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}

	if got, err := parsed.GetTokenID(); err != nil || got != tokenID {
		t.Errorf("token id after Parse: %q, %v", got, err)
	}
	if got, err := parsed.GetIssuer(); err != nil || got != "Google" {
		t.Errorf("issuer after Parse: %q, %v", got, err)
	}
	if got, err := parsed.GetAudience(); err != nil || !slices.Equal(got, []string{"Google", "Facebook"}) {
		t.Errorf("audience after Parse: %v, %v", got, err)
	}
	if got, err := parsed.GetSubject(); err != nil || got != "Billy" {
		t.Errorf("subject after Parse: %q, %v", got, err)
	}
	if got, err := parsed.GetIssuedAt(); err != nil || got != issuedAt.Unix() {
		t.Errorf("issuedAt after Parse: %d, %v", got, err)
	}
	if got, err := parsed.GetNotBeforeAt(); err != nil || got != notBeforeAt.Unix() {
		t.Errorf("notBeforeAt after Parse: %d, %v", got, err)
	}
	if got, err := parsed.GetExpiresAt(); err != nil || got != expiresAt.Unix() {
		t.Errorf("expiresAt after Parse: %d, %v", got, err)
	}
}

// Only a list of strings is an audience. A single string stays an error, as scenario_test.go expects
func TestAudienceInvalid(t *testing.T) {
	values := map[string]any{
		"string":          "Google",
		"number":          5,
		"null":            nil,
		"list of numbers": []any{1, 2},
		"mixed list":      []any{"Google", 1},
		"list of lists":   []any{[]any{"Google"}},
	}

	for name, value := range values {
		claims := New()
		claims.Set(Audience, value)
		audience, err := claims.GetAudience()
		if err != ErrClaimValueInvalid || len(audience) != 0 {
			t.Errorf("%s: got %v, %v, want an empty audience and ErrClaimValueInvalid", name, audience, err)
		}
	}
}

func TestSubject(t *testing.T) {
	claims := New()
	claims.SetSubject("Google")
	subject, _ := claims.GetSubject()
	if subject != "Google" {
		t.Error("subject was not set")
	}
	claims.DeleteSubject()
	if claims.Has(Subject) {
		t.Error("subject should have been deleted")
	}
	subject, _ = claims.GetSubject()
	if subject != "" {
		t.Error("should have gotten blank value")
	}
}

func TestIssuedAt(t *testing.T) {
	now := time.Now()
	claims := New()
	claims.SetIssuedAt(now)
	issuedAt, _ := claims.GetIssuedAt()
	if issuedAt != now.Unix() {
		t.Error("issuedAt was not set")
	}
	claims.DeleteIssuedAt()
	if claims.Has(IssuedAt) {
		t.Error("issuedAt should have been deleted")
	}
	issuedAt, _ = claims.GetIssuedAt()
	if issuedAt != 0 {
		t.Error("should have gotten 0 value")
	}
}
func TestExpiresAt(t *testing.T) {
	now := time.Now()
	claims := New()
	claims.SetExpiresAt(now)
	expiresAt, _ := claims.GetExpiresAt()
	if expiresAt != now.Unix() {
		t.Error("expiresAt was not set")
	}
	claims.DeleteExpiresAt()
	if claims.Has(ExpiresAt) {
		t.Error("expiresAt should have been deleted")
	}
	expiresAt, _ = claims.GetExpiresAt()
	if expiresAt != 0 {
		t.Error("should have gotten 0 value")
	}
}

func TestNotBeforeAt(t *testing.T) {
	now := time.Now()
	claims := New()
	claims.SetNotBeforeAt(now)
	notBeforeAt, _ := claims.GetNotBeforeAt()
	if notBeforeAt != now.Unix() {
		t.Error("notBeforeAt was not set")
	}
	claims.DeleteNotBeforeAt()
	if claims.Has(NotBeforeAt) {
		t.Error("NotBeforeAt should have been deleted")
	}
	notBeforeAt, _ = claims.GetNotBeforeAt()
	if notBeforeAt != 0 {
		t.Error("should have gotten 0 value")
	}
}
