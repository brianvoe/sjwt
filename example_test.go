package sjwt

import (
	"fmt"
	"time"
)

func Example() {
	// Add Claims
	claims := New()
	claims.Add("first_name", "billy")
	claims.Add("last_name", "mister")

	// Generate jwt
	jwt, err := claims.Generate([]byte("secret_key_here"))
	fmt.Println(jwt, err)
}

func Example_registeredClaims() {
	claims := New()
	claims.SetTokenID()
	claims.SetSubject("Subject Title")
	claims.SetIssuer("Google")
	claims.SetAudience([]string{"Google", "Facebook"})
	claims.SetIssuedAt(time.Now())                       // IssuedAt in time, value is set in unix
	claims.SetNotBeforeAt(time.Now().Add(time.Hour * 1)) // Token valid in 1 hour
	claims.SetExpiresAt(time.Now().Add(time.Hour * 24))  // Token expires in 24 hours

	jwt, err := claims.Generate([]byte("secret_key_here"))
	fmt.Println(jwt, err)
}
