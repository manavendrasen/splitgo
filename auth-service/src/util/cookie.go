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
	cookie.Path = "/api/v1"
	c.SetCookie(cookie)
}

func DeleteCookie(c echo.Context, name string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.MaxAge = 0
	c.SetCookie(cookie)
}