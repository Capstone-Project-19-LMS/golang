package courseController

import (
	middlewares "golang/app/middlewares/instructor"
	"golang/constant/constantError"
	"golang/models/dto"
	"golang/service/courseService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CourseController struct {
	CourseService courseService.CourseService
}

// CreateCourse is a function to create course
func (cc *CourseController) CreateCourse(c echo.Context) error {
	var course dto.CourseTransaction
	// Binding request body to struct
	err := c.Bind(&course)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Validate request body
	if err = c.Validate(course); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Get user id from jwt
	claims := middlewares.GetUserInstructor(c)
	course.InstructorID = claims.ID
	// fmt.Println("id = ",course.InstructorID)

	// Call service to create course
	err = cc.CourseService.CreateCourse(course)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail create course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create course",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create course",
	})
}