package middlewares

import (
	"golang/util"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var whitelistCostumer []string = make([]string, 5)

type JwtCustomClaims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(userID uint, role string) (string, error) {
	claims := JwtCustomClaims{
		userID,
		role,
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

func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)

	isListed := CheckToken(user.Raw)

	if !isListed {
		return nil
	}

	claims := user.Claims.(*JwtCustomClaims)
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
