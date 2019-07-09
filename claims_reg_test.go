package sjwt

import (
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
