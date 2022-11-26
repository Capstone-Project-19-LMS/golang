package helper

import (
	middlewareCustomer "golang/app/middlewares/costumer"
	middlewareInstructor "golang/app/middlewares/instructor"
	"golang/models/dto"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) dto.User {
	claims := middlewareInstructor.GetUserInstructor(c)
	var userData dto.User
	if claims != nil {
		userData = dto.User{
			ID: claims.ID,
			Role: claims.Role,
		}
	} else {
		claims := middlewareCustomer.GetUserCustomer(c)
		userData = dto.User{
			ID: claims.ID,
			Role: claims.Role,
		}
	}


	return userData
}