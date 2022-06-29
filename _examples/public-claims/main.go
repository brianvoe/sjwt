package main

import(
	"fmt"
	"github.com/brianvoe/sjwt"
)

func main() {
	// Add Claims
	claims := sjwt.New()
	claims.Set("username", "billymister")
	claims.Set("account_id", 8675309)

	// Generate jwt
	secretKey := []byte("secret_key_here")
	jwt := claims.Generate(secretKey)
	fmt.Println(jwt)
}