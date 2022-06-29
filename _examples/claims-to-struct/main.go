package main

import(
	"fmt"
	"github.com/brianvoe/sjwt"
)

func main(){
	type Info struct {
		Name string `json:"name"`
	}

	// Parse jwt
	jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiQmlsbHkgTWlzdGVyIn0.2FYrpCNy1tg_4UvimpSrgAy-nT9snh-l4w9VLz71b6Y"
	claims, _ := sjwt.Parse(jwt)

	// Marshal your struct into claims
	info := Info{}
	claims.ToStruct(&info)

	name, _ := claims.GetStr("name")
		fmt.Println(name)
	// output: Billy Mister
}