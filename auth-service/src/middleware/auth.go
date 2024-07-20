package middleware

import (
	"auth-service/src/util"
	"log"
	"net/http"

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
		accessTokenString, err := c.Cookie("ACCESS_TOKEN")

		for cc := range c.Cookies() {
			log.Print(c.Cookies()[cc].Value)
		}

		if err != nil || accessTokenString.Value == "" {
			return c.JSON(http.StatusUnauthorized, util.SendMessage("ACCESS_TOKEN_NOT_FOUND"))
		}

		accessToken, err := util.ParseToken(accessTokenString.Value)

		if err != nil {
			log.Print(err)
		}

		claims, ok := accessToken.Claims.(*util.AccessTokenClaim)

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
