package instructorcontroller

import (
	middlewares "golang/app/middlewares/instructor"
	"golang/models/dto"
	instructorService "golang/service/instructorService"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type InstructorController struct {
	InstructorService instructorService.InstructorService
}

func (u *InstructorController) Register(c echo.Context) error {
	var user dto.InstructorRegister
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

	err = u.InstructorService.CreateInstructor(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create instructor",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "success create instructor",
	})
}

func (u *InstructorController) Login(c echo.Context) error {
	var instructorLogin dto.InstructorLogin
	err := c.Bind(&instructorLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}
	var user dto.InstructorResponseGet
	user, err = u.InstructorService.LoginInstructor(instructorLogin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, errToken := middlewares.GenerateTokenInstructor(user.ID)

	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create token",
			"error":   errToken,
		})
	}

	instructorResponse := dto.InstructorResponse{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(200, echo.Map{
		"message": "success login",
		"user":    instructorResponse,
	})
}

func (u *InstructorController) Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckTokenInstructor(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid token",
		})
	}

	middlewares.LogoutInstructor(user.Raw)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "logout success",
	})
}
