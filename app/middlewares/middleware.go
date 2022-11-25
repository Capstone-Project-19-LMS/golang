package middlewares

import (
	customer "golang/app/middlewares/costumer"
	instructor "golang/app/middlewares/instructor"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ConfigLogger struct {
	Format string
}

func (c *ConfigLogger) Init() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: c.Format,
	})
}

func CheckTokenMiddlewareCustomer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customerID := customer.GetUserCustomer(c)

		if customerID == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"messege": "invalid create token",
			})
		}
		return next(c)
	}
}
func CheckTokenMiddlewareInstructor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		InstructorID := instructor.GetUserInstructor(c)

		if InstructorID == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"messege": "invalid create token",
			})
		}
		return next(c)
	}
}
