package sjwt

import (
	"fmt"
	"testing"
)

type benchCase struct {
	name           string
	generateSecret []byte
	verifySecret   []byte
	buildClaims    func() *Claims
	mutateToken    func(string) string
	skipVerify     bool
}

func makeBenchCases() []benchCase {
	return []benchCase{
		{
			name:           "small",
			generateSecret: []byte("benchmark-small-secret-01234567890123"),
			verifySecret:   []byte("benchmark-small-secret-01234567890123"),
			buildClaims: func() *Claims {
				c := New()
				c.Set("user", "123")
				c.Set("role", "admin")
				c.Set("active", true)
				return c
			},
		},
		{
			name:           "medium",
			generateSecret: []byte("benchmark-medium-secret-012345678901"),
			verifySecret:   []byte("benchmark-medium-secret-012345678901"),
			buildClaims: func() *Claims {
				c := New()
				for i := 0; i < 32; i++ {
					c.Set(keyForIndex(i), i)
				}
				return c
			},
		},
		{
			name:           "large",
			generateSecret: []byte("benchmark-large-secret-01234567890123"),
			verifySecret:   []byte("benchmark-large-secret-01234567890123"),
			buildClaims: func() *Claims {
				c := New()
				for i := 0; i < 200; i++ {
					c.Set(keyForIndex(i), i)
				}
				return c
			},
		},
	}
}

func keyForIndex(i int) string {
	return keyPrefix(i) + "_" + fmt.Sprintf("%03d", i)
}

func keyPrefix(i int) string {
	switch {
	case i < 10:
		return "claim_small"
	case i < 100:
		return "claim_medium"
	default:
		return "claim_large"
	}
}

func BenchmarkFullUsage(b *testing.B) {
	cases := makeBenchCases()
	for _, bc := range cases {
		bc := bc
		b.Run(bc.name, func(b *testing.B) {
			claims := bc.buildClaims()
			token, err := claims.Generate(bc.generateSecret)
			if err != nil {
				b.Fatalf("generate failed: %v", err)
			}
			if bc.mutateToken != nil {
				token = bc.mutateToken(token)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				claims := bc.buildClaims()
				token, err := claims.Generate(bc.generateSecret)
				if err != nil {
					b.Fatalf("generate failed: %v", err)
				}

				parsed, err := Parse(token)
				if err != nil {
					b.Fatalf("parse failed: %v", err)
				}

				if !bc.skipVerify && !Verify(token, bc.verifySecret) {
					b.Fatalf("verify failed")
				}

				if err := parsed.Validate(); err != nil {
					b.Fatalf("validate failed: %v", err)
				}
			}
		})
	}
}
