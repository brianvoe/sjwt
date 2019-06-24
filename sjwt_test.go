package sjwt

import (
	"testing"
)

var secretKey = []byte("whats up yall")

func TestGenerate(t *testing.T) {
	claims := New()
	claims.Add("hello", "world")
	jwt, err := claims.Generate(secretKey)
	if err != nil {
		t.Error(err)
	}
	if jwt == "" {
		t.Error("jwt is empty")
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		claims := New()
		claims.Add("hello", "world")
		claims.Generate(secretKey)
	}
}

func TestParseClaims(t *testing.T) {
	claims := New()
	claims.Add("hello", "world")
	jwt, _ := claims.Generate(secretKey)

	newClaims, err := ParseClaims(jwt)
	if err != nil {
		t.Error("error parsing claims")
	}
	if !newClaims.Has("hello") || newClaims.GetStr("hello") != "world" {
		t.Error("error getting claims hello from parsed claims")
	}
}

func TestVerify(t *testing.T) {
	claims := New()
	claims.Add("hello", "world")
	jwt, _ := claims.Generate(secretKey)

	verified := Verify(jwt, secretKey)
	if !verified {
		t.Error("verification failed")
	}

	verified = Verify(jwt, []byte("Bad secret"))
	if verified {
		t.Error("verification should have failed")
	}
}
