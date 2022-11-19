package middlewares

import (
	"golang/util"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var whitelistCostumer []string = make([]string, 5)

type JwtCostumerClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(userID uint) (string, error) {
	claims := JwtCostumerClaims{
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
	whitelistCostumer = append(whitelistCostumer, token)

	return token, nil
}

func GetUser(c echo.Context) *JwtCostumerClaims {
	user := c.Get("user").(*jwt.Token)

	isListed := CheckToken(user.Raw)

	if !isListed {
		return nil
	}

	claims := user.Claims.(*JwtCostumerClaims)
	return claims
}

func CheckToken(token string) bool {
	for _, tkn := range whitelistCostumer {
		if tkn == token {
			return true
		}
	}

	return false
}

func Logout(token string) bool {
	for idx, tkn := range whitelistCostumer {
		if tkn == token {
			whitelistCostumer = append(whitelistCostumer[:idx], whitelistCostumer[idx+1:]...)
		}
	}

	return true
}