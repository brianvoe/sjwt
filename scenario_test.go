package sjwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestUsageScenarios(t *testing.T) {
	type scenario struct {
		name              string
		buildClaims       func() (*Claims, error)
		mutateToken       func(string) string
		secret            []byte
		verifySecret      []byte
		expectGenerateErr error
		expectVerify      bool
		expectParseErr    error
		expectValidateErr error
		skipVerify        bool
		skipParse         bool
		skipValidate      bool
		assertions        func(t *testing.T, claims Claims)
	}

	scenarios := []scenario{
		{
			name: "public claims basic authentication",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.Set("user_id", "123")
				c.Set("role", "admin")
				return c, nil
			},
			secret:       []byte("scenario-public-claims-secret-012345"),
			expectVerify: true,
			assertions: func(t *testing.T, claims Claims) {
				uid, _ := claims.GetStr("user_id")
				if uid != "123" {
					t.Fatalf("expected user_id to be 123, got %s", uid)
				}
			},
		},
		{
			name: "registered claims with validations",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.SetTokenID()
				c.SetIssuer("issuer.example")
				c.SetAudience([]string{"service-a", "service-b"})
				c.SetSubject("user:42")
				c.SetIssuedAt(time.Now().Add(-time.Minute))
				c.SetNotBeforeAt(time.Now().Add(-time.Minute))
				c.SetExpiresAt(time.Now().Add(10 * time.Minute))
				return c, nil
			},
			secret:       []byte("scenario-registered-claims-secret-0123"),
			expectVerify: true,
			assertions: func(t *testing.T, claims Claims) {
				if !claims.Has(TokenID) {
					t.Fatalf("expected token id to be set")
				}
			},
		},
		{
			name: "struct to claims round trip",
			buildClaims: func() (*Claims, error) {
				type profile struct {
					Name  string `json:"name"`
					Email string `json:"email"`
				}
				claims, err := ToClaims(profile{Name: "Billy", Email: "billy@example.com"})
				if err != nil {
					return nil, err
				}
				c := Claims(claims)
				return &c, nil
			},
			secret:       []byte("scenario-struct-round-trip-secret-01"),
			expectVerify: true,
			assertions: func(t *testing.T, claims Claims) {
				type profile struct {
					Name  string `json:"name"`
					Email string `json:"email"`
				}
				var p profile
				if err := claims.ToStruct(&p); err != nil {
					t.Fatalf("failed to convert to struct: %v", err)
				}
				if p.Name != "Billy" {
					t.Fatalf("expected name Billy, got %s", p.Name)
				}
			},
		},
		{
			name: "expired token validation failure",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.Set("user", "legacy")
				c.SetExpiresAt(time.Now().Add(-1 * time.Minute))
				return c, nil
			},
			secret:            []byte("scenario-expired-token-secret-0123"),
			expectVerify:      true,
			expectValidateErr: ErrTokenHasExpired,
		},
		{
			name: "not before validation failure",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.Set("user", "future")
				c.SetNotBeforeAt(time.Now().Add(5 * time.Minute))
				return c, nil
			},
			secret:            []byte("scenario-not-before-secret-0123456"),
			expectVerify:      true,
			expectValidateErr: ErrTokenNotYetValid,
		},
		{
			name: "tampered signature",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.Set("user", "victim")
				return c, nil
			},
			secret:       []byte("scenario-tampered-signature-secret-0"),
			expectVerify: false,
			skipParse:    true,
			mutateToken: func(token string) string {
				parts := strings.Split(token, ".")
				sig, _ := base64.RawURLEncoding.DecodeString(parts[signatureSegmentIdx])
				if len(sig) > 0 {
					sig[0] ^= 0xFF
				}
				parts[signatureSegmentIdx] = base64.RawURLEncoding.EncodeToString(sig)
				return strings.Join(parts, ".")
			},
		},
		{
			name: "algorithm mismatch header",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.Set("user", "header")
				return c, nil
			},
			secret:         []byte("scenario-alg-mismatch-secret-012345"),
			expectVerify:   false,
			expectParseErr: ErrTokenAlgorithmMismatch,
			mutateToken: func(token string) string {
				parts := strings.Split(token, ".")
				parts[headerSegmentIdx] = base64.RawURLEncoding.EncodeToString([]byte(`{"typ":"JWT","alg":"none"}`))
				return strings.Join(parts, ".")
			},
		},
		{
			name: "token missing signature segment",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.Set("user", "missing-sig")
				return c, nil
			},
			secret:         []byte("scenario-missing-signature-secret-012"),
			expectVerify:   false,
			expectParseErr: ErrTokenInvalid,
			mutateToken: func(token string) string {
				parts := strings.Split(token, ".")
				return strings.Join(parts[:2], ".")
			},
		},
		{
			name: "payload tampering without verification",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.Set("order_id", "ABC123")
				c.Set("amount", 130.75)
				return c, nil
			},
			secret:       []byte("scenario-payload-tamper-secret-0123"),
			expectVerify: false,
			assertions: func(t *testing.T, claims Claims) {
				amount, _ := claims.GetFloat("amount")
				if amount != 999.99 {
					t.Fatalf("expected tampered amount 999.99, got %v", amount)
				}
			},
			mutateToken: func(token string) string {
				parts := strings.Split(token, ".")
				payload, err := base64.RawURLEncoding.DecodeString(parts[payloadSegmentIdx])
				if err != nil {
					return token
				}
				var data map[string]any
				if err := json.Unmarshal(payload, &data); err != nil {
					return token
				}
				data["amount"] = 999.99
				b, err := json.Marshal(data)
				if err != nil {
					return token
				}
				parts[payloadSegmentIdx] = base64.RawURLEncoding.EncodeToString(b)
				return strings.Join(parts, ".")
			},
		},
		{
			name: "large payload with many claims",
			buildClaims: func() (*Claims, error) {
				c := New()
				for i := 0; i < 200; i++ {
					c.Set(fmt.Sprintf("key_%03d", i), i)
				}
				return c, nil
			},
			secret:       []byte("scenario-large-payload-secret-01234"),
			expectVerify: true,
			assertions: func(t *testing.T, claims Claims) {
				if len(claims) < 200 {
					t.Fatalf("expected at least 200 claims, got %d", len(claims))
				}
			},
		},
		{
			name: "audience type mismatch detection",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.Set(Audience, "single-audience")
				return c, nil
			},
			secret:       []byte("scenario-audience-mismatch-secret-0123"),
			expectVerify: true,
			skipValidate: true,
			assertions: func(t *testing.T, claims Claims) {
				if _, err := claims.GetAudience(); err != ErrClaimValueInvalid {
					t.Fatalf("expected ErrClaimValueInvalid, got %v", err)
				}
			},
		},
		{
			name: "rotated secret verification fails",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.Set("session", "original")
				return c, nil
			},
			secret:       []byte("scenario-rotated-secret-original-0123"),
			verifySecret: []byte("scenario-rotated-secret-new-012345"),
			expectVerify: false,
			assertions: func(t *testing.T, claims Claims) {
				session, _ := claims.GetStr("session")
				if session != "original" {
					t.Fatalf("expected session original, got %s", session)
				}
			},
		},
		{
			name: "short secret generation failure",
			buildClaims: func() (*Claims, error) {
				c := New()
				c.Set("mode", "insecure")
				return c, nil
			},
			secret:            []byte("short secret"),
			expectGenerateErr: ErrSecretTooShort,
			skipVerify:        true,
			skipParse:         true,
		},
	}

	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			secret := sc.secret
			if len(secret) < minSecretLength {
				if sc.expectGenerateErr != nil {
					// Allow short secret when we expect generation to fail.
				} else {
					t.Fatalf("scenario %s secret must be at least %d bytes", sc.name, minSecretLength)
				}
			}

			claims, err := sc.buildClaims()
			if err != nil {
				t.Fatalf("failed to build claims: %v", err)
			}

			token, err := claims.Generate(secret)
			if sc.expectGenerateErr != nil {
				if !errors.Is(err, sc.expectGenerateErr) {
					t.Fatalf("expected generate error %v, got %v", sc.expectGenerateErr, err)
				}
				if err == nil {
					t.Fatalf("expected generate error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("failed to generate token: %v", err)
			}

			if sc.mutateToken != nil {
				token = sc.mutateToken(token)
			}

			if !sc.skipVerify {
				verifySecret := sc.verifySecret
				if verifySecret == nil {
					verifySecret = secret
				}
				if verified := Verify(token, verifySecret); verified != sc.expectVerify {
					t.Fatalf("expected verify to be %t, got %t", sc.expectVerify, verified)
				}
			}

			if sc.skipParse {
				return
			}

			parsedClaims, err := Parse(token)
			if sc.expectParseErr != nil {
				if !errors.Is(err, sc.expectParseErr) {
					t.Fatalf("expected parse error %v, got %v", sc.expectParseErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected parse error: %v", err)
			}

			if sc.skipValidate {
				if sc.assertions != nil {
					sc.assertions(t, parsedClaims)
				}
				return
			}

			if sc.expectValidateErr != nil {
				if err := parsedClaims.Validate(); !errors.Is(err, sc.expectValidateErr) {
					t.Fatalf("expected validate error %v, got %v", sc.expectValidateErr, err)
				}
			} else if err := parsedClaims.Validate(); err != nil {
				t.Fatalf("unexpected validation error: %v", err)
			}

			if sc.assertions != nil {
				sc.assertions(t, parsedClaims)
			}
		})
	}
}
