package util

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func SetTokenInCookie(c echo.Context, claims jwt.MapClaims, name string, expirationTime time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		return "", errors.New("FAILED_TOKEN_GENERATION")
	}

	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = tokenString
	cookie.Expires = expirationTime
	cookie.HttpOnly = true
	cookie.Secure = true
	c.SetCookie(cookie)

	return tokenString, nil
}

func GetTokenInCookie(c echo.Context, name string) (*jwt.Token, string, error) {

	tokenString, err := c.Cookie(name)
	if err != nil {
		return nil, "", errors.New("ERROR_GETTING_COOKIE")
	}

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := ParseToken(tokenString.Value)

	if err != nil {
		return nil, "", errors.New("ERROR_GETTING_COOKIE")
	}
	return token, tokenString.Value, nil
}

func ParseToken (tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_KEY")), nil
	})
}
