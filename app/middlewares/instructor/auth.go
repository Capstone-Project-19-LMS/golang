package middlewares

import (
	"golang/util"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var whitelistInstructor []string = make([]string, 5)

type JwtInstructorClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
	claims := JwtInstructorClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 2).Unix(),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(util.GetConfig("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}
	whitelistInstructor = append(whitelistInstructor, token)

	return token, nil
}

func GetUser(c echo.Context) *JwtInstructorClaims {
	user := c.Get("user").(*jwt.Token)

	isListed := CheckToken(user.Raw)

	if !isListed {
		return nil
	}

	claims := user.Claims.(*JwtInstructorClaims)
	return claims
}

func CheckToken(token string) bool {
	for _, tkn := range whitelistInstructor {
		if tkn == token {
			return true
		}
	}

	return false
}

func Logout(token string) bool {
	for idx, tkn := range whitelistInstructor {
		if tkn == token {
			whitelistInstructor = append(whitelistInstructor[:idx], whitelistInstructor[idx+1:]...)
		}
	}

	return true
}
