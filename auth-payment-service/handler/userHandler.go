package handler

import (
	"net/http"
	"os"
	"payment-service/model"
	"payment-service/repository"
	"payment-service/util"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c echo.Context) error {
	var body struct {
		Email          string
		Password       string
		PhoneNumber    string
		DisplayName    string
		ProfilePicture string
	}

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadGateway, util.SendMessage("INVALID_BODY"))
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		return c.JSON(http.StatusBadGateway, util.SendMessage("FAILED_HASH"))
	}

	user := model.User{
		DisplayName:    body.DisplayName,
		Email:          body.Email,
		PhoneNumber:    body.PhoneNumber,
		Password:       string(hash),
		ProfilePicture: body.ProfilePicture,
	}

	err = repository.SignUp(&user)

	if err != nil {
		return c.JSON(http.StatusBadGateway, util.SendMessage("DB_INSERT_ERROR"))
	}

	return c.JSON(http.StatusOK, util.SendMessage("SUCCESS"))
}

func Login(c echo.Context) error {
	var body struct {
		Email    string
		Password string
	}

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadGateway, util.SendMessage("INVALID_BODY"))
	}

	user, err := repository.Login(body.Email)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, util.SendMessage(err.Error()))
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, util.SendMessage("INVALID_EMAIL_OR_PASSWORD"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":          uint(user.ID),
		"DisplayName": string(user.DisplayName),
		"Email":       string(user.Email),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		return c.JSON(http.StatusBadGateway, util.SendMessage("FAILED_TOKEN_VERIFICATION"))
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, util.SendMessage("SUCCESS"))
}
