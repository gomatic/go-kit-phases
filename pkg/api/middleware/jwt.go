package middleware

import (
	"crypto/rand"
	"encoding/base32"
	"os"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	stdjwt "github.com/golang-jwt/jwt/v4"
)

func JWT(c endpoint.Endpoint) endpoint.Endpoint {
	// TODO make configurable (generalize with other endpoints)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		randomBytes := make([]byte, 32)
		_, err := rand.Read(randomBytes)
		if err != nil {
			panic(err)
		}
		secret = base32.StdEncoding.EncodeToString(randomBytes)
	}
	kf := func(token *stdjwt.Token) (interface{}, error) { return []byte(secret), nil }
	return jwt.NewParser(kf, stdjwt.SigningMethodHS256, jwt.StandardClaimsFactory)(c)
}
