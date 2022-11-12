package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/urfave/cli/v2"
)

//
func action(c *cli.Context) error {
	args := c.Args()

	now := time.Now().UTC()
	u := now.Unix()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"nbf": u,
			"iat": u,
			"exp": now.Add(24 * time.Hour).Unix(),
		},
	)

	secret := args.First()
	if secret == "" {
		secret = c.String("secret")
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	fmt.Println(tokenString)
	return nil
}

//
func Commands() cli.Commands {
	return cli.Commands{
		{
			Name:        "token",
			Usage:       "client test token [JWT-SECRET]",
			Description: "client test token [JWT-SECRET]",
			Action:      action,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "secret",
					EnvVars: []string{"JWT_SECRET"},
				},
			},
		},
	}
}
