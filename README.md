![alt text](https://raw.githubusercontent.com/brianvoe/sjwt/master/logo.png)

# sjwt [![Go Report Card](https://goreportcard.com/badge/github.com/brianvoe/sjwt)](https://goreportcard.com/report/github.com/brianvoe/sjwt) [![GoDoc](https://godoc.org/github.com/brianvoe/sjwt?status.svg)](https://godoc.org/github.com/brianvoe/sjwt) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/brianvoe/sjwt/master/LICENSE)

<a href="https://www.buymeacoffee.com/brianvoe" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>
## Install

`go get -u github.com/brianvoe/sjwt`

Simple JSON Web Token - Uses HMAC SHA-256

Minimalistic and efficient tool for handling JSON Web Tokens in Go applications. It offers a straightforward approach to integrating JWT for authentication and security, designed for ease of use.

## Features

- **Easy JWT for Go**: Implement JWT in Go with minimal effort.
- **Secure & Simple**: Reliable security features, easy to integrate.
- **Open Source**: MIT licensed, open for community contributions.

## Install
```bash 
go get -u github.com/brianvoe/sjwt
```

## Example
```go
// Set Claims
claims := sjwt.New()
claims.Set("username", "billymister")
claims.Set("account_id", 8675309)

// Generate jwt
secretKey := []byte("secret_key_here")
jwt := claims.Generate(secretKey)
```

## Example parse
```go
// Parse jwt
jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
claims, _ := sjwt.Parse(jwt)

// Get claims
name, err := claims.GetStr("name") // John Doe
```

## Example verify and validate
```go
jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
secretKey := []byte("secret_key_here")

// Verify that the secret signature is valid
hasVerified := sjwt.Verify(jwt, secretKey)

// Parse jwt
claims, _ := sjwt.Parse(jwt)

// Validate will check(if set) Expiration At and Not Before At dates
err := claims.Validate()
```

## Example usage of registered claims
```go
// Set Claims
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
```

## Example usage of struct to claims
```go
type Info struct {
    Name string `json:"name"`
}

// Marshal your struct into claims
info := Info{Name: "Billy Mister"}
claims, _ := sjwt.ToClaims(info)

// Generate jwt
secretKey := []byte("secret_key_here")
jwt := claims.Generate(secretKey)
```

## Why?
For all the times I have needed the use of a jwt, its always been a simple HMAC SHA-256 and thats normally the use of most jwt tokens.
