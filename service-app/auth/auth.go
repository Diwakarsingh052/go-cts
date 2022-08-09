package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type ctxKey int

// Claims is our payload/data for out jwt token
type Claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

const (
	RoleAdmin = "ADMIN"
	RoleUser  = "USER"
)

func (c Claims) HasRoles(roles ...string) bool {

	for _, has := range c.Roles { // roles with the user
		for _, want := range roles { // what roles  handler demand
			if has == want {
				return true
			}
		}
	}

	return false

}

// Key is used to store/retrieve a Claims value from a context.Context.
const Key ctxKey = 1

// Auth struct privateKey field would be used to verify and generate token
type Auth struct {
	privateKey *rsa.PrivateKey // this is key we would get after parsing the private.pem file
}

// NewAuth func set the privateKey in the Auth struct and returns the instance of it to the caller
func NewAuth(privateKey *rsa.PrivateKey) (*Auth, error) {
	if privateKey == nil {
		return nil, errors.New("private key cannot be nil")
	}

	a := Auth{privateKey: privateKey}
	return &a, nil

}

func (a *Auth) GenerateToken(claims Claims) (string, error) {

	//jwt.NewWithClaims takes a signingMethod and claims struct to generate a token on basis of it
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, &claims)

	//signing our token with our private key
	tokenStr, err := tkn.SignedString(a.privateKey)

	if err != nil {
		return "", fmt.Errorf("signing token %w", err)
	}

	return tokenStr, nil

}

func (a *Auth) ValidateToken(tokenStr string) (Claims, error) {

	var claims Claims

	// verifying token with our public key
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return a.privateKey.Public(), nil
	})

	if err != nil {
		if !token.Valid {
			return Claims{}, errors.New("invalid token")
		}
		return Claims{}, fmt.Errorf("parsing token %w", err)
	}

	return claims, nil

}
