package sjwt

import (
	"fmt"
	"time"
)

func Example() {
	// Add Claims
	claims := New()
	claims.Set("username", "billymister")
	claims.Set("account_id", 8675309)

	// Generate jwt
	secretKey := []byte("secret_key_here")
	jwt := claims.Generate(secretKey)
	fmt.Println(jwt)
}

func Example_parse() {
	// Parse jwt
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	claims, _ := Parse(jwt)

	// Get claims
	name, _ := claims.GetStr("name")
	fmt.Println(name)
	// Output: John Doe
}

func Example_registeredClaims() {
	// Add Claims
	claims := New()
	claims.SetTokenID()                                  // UUID generated
	claims.SetSubject("Subject Title")                   // Subject of the token
	claims.SetIssuer("Google")                           // Issuer of the token
	claims.SetAudience([]string{"Google", "Facebook"})   // Audience the toke is for
	claims.SetIssuedAt(time.Now())                       // IssuedAt in time, value is set in unix
	claims.SetNotBeforeAt(time.Now().Add(time.Hour * 1)) // Token valid in 1 hour
	claims.SetExpiresAt(time.Now().Add(time.Hour * 24))  // Token expires in 24 hours

	// Generate jwt
	secretKey := []byte("secret_key_here")
	jwt := claims.Generate(secretKey)
	fmt.Println(jwt)
}

func Example_publicClaims() {
	// Add Claims
	claims := New()
	claims.Set("username", "billymister")
	claims.Set("account_id", 8675309)

	// Generate jwt
	secretKey := []byte("secret_key_here")
	jwt := claims.Generate(secretKey)
	fmt.Println(jwt)
}

func Example_structToClaims() {
	type Info struct {
		Name string `json:"name"`
	}

	// Marshal your struct into claims
	info := Info{Name: "Billy Mister"}
	claims, _ := ToClaims(info)

	// Generate jwt
	secretKey := []byte("secret_key_here")
	jwt := claims.Generate(secretKey)
	fmt.Println(jwt)
	// output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQmlsbHkgTWlzdGVyIn0.2FYrpCNy1tg_4UvimpSrgAy-nT9snh-l4w9VLz71b6Y
}

func Example_claimsToStruct() {
	type Info struct {
		Name string `json:"name"`
	}

	// Parse jwt
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQmlsbHkgTWlzdGVyIn0.2FYrpCNy1tg_4UvimpSrgAy-nT9snh-l4w9VLz71b6Y"
	claims, _ := Parse(jwt)

	// Marshal your struct into claims
	info := Info{}
	claims.ToStruct(&info)

	name, _ := claims.GetStr("name")
	fmt.Println(name)
	// output: Billy Mister
}

func Example_verifySignature() {
	secretKey := []byte("secret_key_here")
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQmlsbHkgTWlzdGVyIn0.2FYrpCNy1tg_4UvimpSrgAy-nT9snh-l4w9VLz71b6Y"

	// Pass jwt and secret key to verify
	verified := Verify(jwt, secretKey)
	fmt.Println(verified)
	// output: true
}
