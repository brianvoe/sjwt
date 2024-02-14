package main

import(
	"fmt"
	"github.com/brianvoe/sjwt"
)

func main() {
	type Info struct {
		Name string `json:"name"`
	}

	// Marshal your struct into claims
	info := Info{Name: "Billy Mister"}
	claims, _ := sjwt.ToClaims(info)

	// Generate jwt
	secretKey := []byte("secret_key_here")
	jwt := claims.Generate(secretKey)
	fmt.Println(jwt)
	// output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQmlsbHkgTWlzdGVyIn0.2FYrpCNy1tg_4UvimpSrgAy-nT9snh-l4w9VLz71b6Y
}