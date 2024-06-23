package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthContext struct {
	echo.Context
	userID          uint
	userEmail       string
	userDisplayName string
}

func (c *AuthContext) GetCurrentUser() (uint, string, string) {
	return c.userID, c.userEmail, c.userDisplayName
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		tokenString, err := c.Cookie("auth")
		if err != nil {
			return err
		}

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			log.Fatal(err)
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if ok {
			contextWithUserDetails := &AuthContext{
				Context:         c,
				userID:          uint(claims["ID"].(float64)),
				userEmail:       claims["Email"].(string),
				userDisplayName: claims["DisplayName"].(string),
			}
			return next(contextWithUserDetails)
		} else {
			log.Fatal("Error Getting Cookie")
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
}
