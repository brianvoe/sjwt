package main

import(
	"fmt"
	"github.com/brianvoe/sjwt"
)

func main() {
	secretKey := []byte("secret_key_here")
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQmlsbHkgTWlzdGVyIn0.2FYrpCNy1tg_4UvimpSrgAy-nT9snh-l4w9VLz71b6Y"

	// Pass jwt and secret key to verify
	verified := sjwt.Verify(jwt, secretKey)
	fmt.Println(verified)
	// output: true
}