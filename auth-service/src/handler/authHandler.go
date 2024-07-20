package handler

import (
	"auth-service/src/database"
	"auth-service/src/model"
	"auth-service/src/repository"
	"auth-service/src/util"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func getAccessTokenExpirationTime() (time.Time) {
	return time.Now().Add(10 * time.Minute)
}
func getRefreshTokenExpirationTime() (time.Time) {
	return time.Now().Add(5 * time.Hour)
}

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

	// get user details
	user, err := repository.Login(body.Email)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, util.SendMessage(err.Error()))
	}

	// check password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, util.SendMessage("INVALID_EMAIL_OR_PASSWORD"))
	}

	return createAndHandleTokens(c, user)
}

func Refresh(c echo.Context) error {

	// get token strings from cookies
	accessTokenString, err := c.Cookie("ACCESS_TOKEN")

	if err != nil {
		return c.JSON(http.StatusForbidden, util.SendMessage("INVALID_TOKEN 1"))
	}

	refreshTokenString, err := c.Cookie("REFRESH_TOKEN")

	if err != nil {
		return c.JSON(http.StatusForbidden, util.SendMessage("INVALID_TOKEN 2"))
	}

	storedRefreshTokenString, err := database.GetCache(accessTokenString.Value)

	// delete cache

	if err != nil {
		return c.JSON(http.StatusForbidden, util.SendMessage("INVALID_TOKEN 3"))
	}

	if strings.Compare(storedRefreshTokenString, refreshTokenString.Value) != 0 {
		return c.JSON(http.StatusBadGateway, util.SendMessage("INVALID_REFRESH_TOKEN"))
	}

	// check if access token is expired
	accessToken, _ := util.ParseToken(accessTokenString.Value)
	refreshToken, _ := util.ParseToken(refreshTokenString.Value)

	accessTokenExpirationTime, err := accessToken.Claims.GetExpirationTime()
	if err != nil {
		// show access forbidden
		return c.JSON(http.StatusForbidden, util.SendMessage("INVALID_TOKEN 4"))
	}

	if time.Now().Before(accessTokenExpirationTime.Time) {
		// auth token is not expired
		return c.JSON(http.StatusOK, util.SendMessage("VALID"))
	}

	// check if refresh Token is expired
	refreshTokenExpirationTime, err := refreshToken.Claims.GetExpirationTime()
	if err != nil {
		// show access forbidden
		return c.JSON(http.StatusForbidden, util.SendMessage("INVALID_TOKEN 5"))
	}

	if time.Now().After(refreshTokenExpirationTime.Time) {
		// refresh token is expired
		return c.JSON(http.StatusForbidden, util.SendMessage("REFRESH_TOKEN_EXPIRED"))
	}

	database.DeleteCache(accessTokenString.Value)
 
	log.Print(accessToken.Claims)
	// auth token is expired and refresh token is not expired
	claims, ok := accessToken.Claims.(*util.AccessTokenClaim)

	if !ok {
		return c.JSON(http.StatusBadGateway, util.SendMessage("INVALID_TOKEN 6"))
	}

	user, err := repository.GetUserByEmail(claims.Email)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, util.SendMessage(err.Error()))
	}

	return createAndHandleTokens(c, user)
}

func createAndHandleTokens(c echo.Context, user *model.User) error {
	// create access token
	accessTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, util.AccessTokenClaim{
		ID:          uint(user.ID),
		DisplayName: string(user.DisplayName),
		Email:       string(user.Email),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(getAccessTokenExpirationTime()),
		},
	}).SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		return errors.New("FAILED_TOKEN_GENERATION")
	}

	// create refresh token
	refreshTokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(getRefreshTokenExpirationTime()),
	}).SignedString([]byte(os.Getenv("JWT_KEY")))

	if err != nil {
		return c.JSON(http.StatusBadGateway, util.SendMessage(err.Error()))
	}

	// set HTTP-ONLY Cookies
	util.SetCookie(c, "ACCESS_TOKEN", accessTokenString)
	util.SetCookie(c, "REFRESH_TOKEN", refreshTokenString)

	// set the access token and refresh token pair in cache
	
	err = database.SetCache(accessTokenString, refreshTokenString)

	if err != nil {
		return c.JSON(http.StatusBadGateway, util.SendMessage("COULD_NOT_SET_TOKENS"))
	}

	return c.JSON(http.StatusOK, util.SendMessage("SUCCESS"))
}
