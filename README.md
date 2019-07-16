![alt text](https://raw.githubusercontent.com/brianvoe/sjwt/master/logo.png)

# sjwt [![Go Report Card](https://goreportcard.com/badge/github.com/brianvoe/sjwt)](https://goreportcard.com/report/github.com/brianvoe/sjwt) [![Build Status](https://travis-ci.org/brianvoe/sjwt.svg?branch=master)](https://travis-ci.org/brianvoe/sjwt) [![codecov.io](https://codecov.io/github/brianvoe/sjwt/branch/master/graph/badge.svg)](https://codecov.io/github/brianvoe/sjwt) [![GoDoc](https://godoc.org/github.com/brianvoe/sjwt?status.svg)](https://godoc.org/github.com/brianvoe/sjwt) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/brianvoe/sjwt/master/LICENSE)
Simple JSON Web Token - Uses HMAC SHA-256

<a href="https://www.buymeacoffee.com/brianvoe" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

## Example
```go
// Set Claims
claims := New()
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
claims, _ := Parse(jwt)

// Get claims
name, err := claims.GetStr("name") // John Doe
```

## Example verify and validate
```go
jwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
secretKey := []byte("secret_key_here")

// Verify that the secret signature is valid
hasVerified := Verify(jwt, secretKey)

// Parse jwt
claims, _ := Parse(jwt)

// Validate will check(if set) Expiration At and Not Before At dates
err := claims.Validate()
```

## Example usage of registered claims
```go
// Set Claims
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
```

## Example usage of struct to claims
```go
type Info struct {
    Name string `json:"name"`
}

// Marshal your struct into claims
info := Info{Name: "Billy Mister"}
claims, _ := ToClaims(info)

// Generate jwt
secretKey := []byte("secret_key_here")
jwt := claims.Generate(secretKey)
```

## Why?
For all the times I have needed the use of a jwt, its always been a simple HMAC SHA-256 and thats normally the use of most jwt tokens.
