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
	secretKey := []byte("0123456789abcdef0123456789abcdef")
	jwt, err := claims.Generate(secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}

func Example_parse() {
	// Create jwt
	secretKey := []byte("0123456789abcdef0123456789abcdef")
	claims := New()
	claims.Set("name", "John Doe")

	jwt, err := claims.Generate(secretKey)
	if err != nil {
		panic(err)
	}

	// Parse jwt (signature still needs to be verified separately)
	parsed, err := Parse(jwt)
	if err != nil {
		panic(err)
	}

	// Get claims
	name, _ := parsed.GetStr("name")
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
	secretKey := []byte("0123456789abcdef0123456789abcdef")
	jwt, err := claims.Generate(secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}

func Example_publicClaims() {
	// Add Claims
	claims := New()
	claims.Set("username", "billymister")
	claims.Set("account_id", 8675309)

	// Generate jwt
	secretKey := []byte("0123456789abcdef0123456789abcdef")
	jwt, err := claims.Generate(secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(jwt)
}

func Example_structToClaims() {
	type Info struct {
		Name string `json:"name"`
	}

	// Marshal your struct into claims
	info := Info{Name: "Billy Mister"}
	claims, err := ToClaims(info)
	if err != nil {
		panic(err)
	}

	// Generate jwt
	secretKey := []byte("0123456789abcdef0123456789abcdef")
	jwt, err := claims.Generate(secretKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(jwt) > 0)
	// output: true
}

func Example_claimsToStruct() {
	type Info struct {
		Name string `json:"name"`
	}

	// Create jwt
	secretKey := []byte("0123456789abcdef0123456789abcdef")
	claims := New()
	claims.Set("name", "Billy Mister")

	jwt, err := claims.Generate(secretKey)
	if err != nil {
		panic(err)
	}

	// Parse jwt
	parsed, err := Parse(jwt)
	if err != nil {
		panic(err)
	}

	// Marshal your struct into claims
	info := Info{}
	if err := parsed.ToStruct(&info); err != nil {
		panic(err)
	}

	name, _ := parsed.GetStr("name")
	fmt.Println(name)
	// output: Billy Mister
}

func Example_verifySignature() {
	secretKey := []byte("0123456789abcdef0123456789abcdef")
	claims := New()
	claims.Set("name", "Billy Mister")

	token, err := claims.Generate(secretKey)
	if err != nil {
		panic(err)
	}

	verified := Verify(token, secretKey)
	fmt.Println(verified)
	// output: true
}

func Example_parse_roundTrip() {
	secretKey := []byte("0123456789abcdef0123456789abcdef")
	claims := New()
	claims.Set("name", "Billy Mister")

	token, err := claims.Generate(secretKey)
	if err != nil {
		panic(err)
	}

	parsedClaims, err := Parse(token)
	if err != nil {
		panic(err)
	}

	fmt.Println(parsedClaims.Has("name"))
	// output: true
}
