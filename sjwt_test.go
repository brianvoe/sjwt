package sjwt

import (
	"encoding/base64"
	"testing"
)

var secretKey = []byte("0123456789abcdef0123456789abcdef")

func TestGenerate(t *testing.T) {
	claims := New()
	claims.Set("hello", "world")
	jwt, err := claims.Generate(secretKey)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}
	if jwt == "" {
		t.Error("jwt is empty")
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		claims := New()
		claims.Set("hello", "world")
		if _, err := claims.Generate(secretKey); err != nil {
			b.Fatalf("Generate returned error: %v", err)
		}
	}
}

func TestParse(t *testing.T) {
	claims := New()
	claims.Set("hello", "world")
	jwt, err := claims.Generate(secretKey)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	newClaims, err := Parse(jwt)
	if err != nil {
		t.Error("error parsing claims")
	}
	if !newClaims.Has("hello") {
		t.Error("error getting claims hello from parsed claims")
	}

	hello, _ := newClaims.GetStr("hello")
	if hello != "world" {
		t.Error("error hello does not equal world")
	}
}

func TestParseEmpty(t *testing.T) {
	_, err := Parse("")
	if err != ErrTokenInvalid {
		t.Error("error should have failed to parse empty jwt")
	}
}

func TestParseDecodeError(t *testing.T) {
	_, err := Parse("..")
	if err == nil {
		t.Error("error should have failed to parse empty jwt")
	}
}

func TestVerify(t *testing.T) {
	claims := New()
	claims.Set("hello", "world")
	jwt, err := claims.Generate(secretKey)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	verified := Verify(jwt, secretKey)
	if !verified {
		t.Error("verification failed")
	}

	verified = Verify(jwt, []byte("Bad secret"))
	if verified {
		t.Error("verification should have failed")
	}

	verified = Verify(jwt, []byte("short secret"))
	if verified {
		t.Error("verification should have failed with short secret")
	}
}

func TestVerifyError(t *testing.T) {
	jwt := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9." +
		"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ." +
		"uk1qJnGuGHHGFw6fXpVILrdo52JqyD3EzvW3_DxhgZPAqU-OKzzPy7xdRNeQRba5CI6VGmlo6DBYqRCteiiOTw"

	verified := Verify(jwt, secretKey)
	if verified {
		t.Error("verification should have failed")
	}
}

func TestVerifyInvalidJWTError(t *testing.T) {
	jwt := "not_a_jwt"

	verified := Verify(jwt, secretKey)
	if verified {
		t.Error("verification should have failed")
	}
}

func TestGenerateSecretTooShort(t *testing.T) {
	claims := New()
	claims.Set("hello", "world")
	if _, err := claims.Generate([]byte("short secret")); err != ErrSecretTooShort {
		t.Fatalf("expected ErrSecretTooShort, got %v", err)
	}
}

func TestParseAlgorithmMismatch(t *testing.T) {
	// Header with alg none
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"typ":"JWT","alg":"none"}`))
	payload := base64.RawURLEncoding.EncodeToString([]byte(`{"sub":"1234567890"}`))
	token := header + "." + payload + "."
	if _, err := Parse(token); err != ErrTokenAlgorithmMismatch {
		t.Fatalf("expected ErrTokenAlgorithmMismatch, got %v", err)
	}
}
