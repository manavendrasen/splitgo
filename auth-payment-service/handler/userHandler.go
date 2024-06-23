package handler

import (
	"net/http"
	"os"
	"payment-service/model"
	"payment-service/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Status struct {
	Message string
}

func SignUp(c echo.Context) error {
	var body struct {
		Email       string
		Password    string
		PhoneNumber string
		DisplayName string
	}

	err := c.Bind(&body)
	if err != nil {
		status := &Status{
			Message: "INVALID_BODY",
		}
		return c.JSON(http.StatusBadGateway, status)
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		status := &Status{
			Message: "FAILED_HASH",
		}
		return c.JSON(http.StatusBadGateway, status)
	}

	user := model.User{
		DisplayName:    body.DisplayName,
		Email:          body.Email,
		PhoneNumber:    body.PhoneNumber,
		Password:       string(hash),
		ProfilePicture: "",
	}

	err = repository.SignUp(&user)

	if err != nil {
		status := &Status{
			Message: "DB_INSERT_ERROR",
		}
		return c.JSON(http.StatusBadGateway, status)
	}

	status := &Status{
		Message: "SUCCESS",
	}
	return c.JSON(http.StatusOK, status)
}

func Login(c echo.Context) error {
	var body struct {
		Email    string
		Password string
	}

	err := c.Bind(&body)
	if err != nil {
		status := &Status{
			Message: "INVALID_BODY",
		}
		return c.JSON(http.StatusBadGateway, status)
	}

	user, err := repository.Login(body.Email)

	if err != nil {
		status := &Status{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, status)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		status := &Status{
			Message: "INVALID_EMAIL_OR_PASSWORD",
		}
		return c.JSON(http.StatusUnauthorized, status)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":          uint(user.ID),
		"DisplayName": string(user.DisplayName),
		"Email":       string(user.Email),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		status := &Status{
			Message: "FAILED_TOKEN",
		}
		return c.JSON(http.StatusBadGateway, status)
	}

	cookie := new(http.Cookie)
	cookie.Name = "auth"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = true
	c.SetCookie(cookie)

	status := &Status{
		Message: "SUCCESS",
	}
	return c.JSON(http.StatusOK, status)
}
