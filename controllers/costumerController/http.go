package CostumerController

import (
	middlewares "golang/app/middlewares/costumer"
	"golang/models/dto"
	costumerService "golang/service/costumerService"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CostumerController struct {
	CostumerService costumerService.CostumerService
}

func (u *CostumerController) Register(c echo.Context) error {
	var user dto.CostumerRegister
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(500, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// validate data user
	err = c.Validate(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	err = u.CostumerService.CreateCustomer(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create user",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "success create user",
	})
}

func (u *CostumerController) Login(c echo.Context) error {
	var costumerLogin dto.CostumerLogin
	err := c.Bind(&costumerLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}
	var user dto.CostumerResponseGet
	user, err = u.CostumerService.LoginCostumer(costumerLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, errToken := middlewares.GenerateToken(user.ID)

	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create token",
			"error":   errToken,
		})
	}

	costumerResponse := dto.CostumerResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(200, echo.Map{
		"message": "success login",
		"user":    costumerResponse,
	})
}

func (u *CostumerController) Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid token",
		})
	}

	middlewares.Logout(user.Raw)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "logout success",
	})
}
