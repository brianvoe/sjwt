package main

import(
	"fmt"
	"github.com/brianvoe/sjwt"
)

func main() {
	// Add Claims
	claims := sjwt.New()
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