package middleware

import (
	"auth-service/src/util"
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

		tokenString, err := c.Cookie("ACCESS_TOKEN")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, util.SendMessage("ACCESS_TOKEN_NOT_FOUND"))
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
			log.Print(err)
		}

		claims, ok := token.Claims.(util.AccessTokenClaim)

		if ok {
			contextWithUserDetails := &AuthContext{
				Context:         c,
				userID:          uint(claims.ID),
				userEmail:       claims.Email,
				userDisplayName: claims.DisplayName,
			}
			return next(contextWithUserDetails)
		} else {
			c.JSON(http.StatusUnauthorized, util.SendMessage("INVALID_ACCESS_TOKEN"))
		}
		return c.JSON(http.StatusUnauthorized, util.SendMessage("ACCESS_TOKEN_NOT_FOUND"))
	}
}
