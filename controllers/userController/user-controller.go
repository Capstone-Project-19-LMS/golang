package userController

import (
	"golang/app/middlewares"
	"golang/models/dto"
	"golang/service/userService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService userService.UserService
}

func (u *UserController) Register(c echo.Context) error {
	var user dto.UserRegister
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

	err = u.UserService.CreateUser(user)
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

func (u *UserController) Login(c echo.Context) error {
	var userLogin dto.UserLogin
	err := c.Bind(&userLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}
	var user dto.UserResponseGet
	user, err = u.UserService.LoginUser(userLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail login",
			"error":   err.Error(),
		})
	}
	
	token, errToken := middlewares.GenerateToken(user.ID, user.Role)

	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create token",
			"error":   errToken,
		})
	}

	userResponse := dto.UserResponse{
		Name: user.Name,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(200, echo.Map{
		"message": "success login",
		"user":   userResponse,
	})
}

func (u *UserController) Logout(c echo.Context) error {
	return nil
}