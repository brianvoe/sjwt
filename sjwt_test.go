package sjwt

import (
	"testing"
)

var secretKey = ""

func TestGenerate(t *testing.T) {
	claims := New()
	claims.AddClaim("hello", "world")
	jwt, err := claims.Generate([]byte("whats up yall"))
	if err != nil {
		t.Error(err)
	}
	if jwt == "" {
		t.Error("jwt is empty")
	}
}
