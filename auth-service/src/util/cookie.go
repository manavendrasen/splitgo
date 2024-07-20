package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetCookie(c echo.Context, name string, tokenString string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = tokenString
	cookie.HttpOnly = true
	cookie.Secure = true
	c.SetCookie(cookie)
}
