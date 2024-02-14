package main

import(
	"fmt"
	"github.com/brianvoe/sjwt"
)

func main() {
	// Parse jwt
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	claims, _ := sjwt.Parse(jwt)

	// Get claims
	name, _ := claims.GetStr("name")
	fmt.Println(name)
	// Output: John Doe
}