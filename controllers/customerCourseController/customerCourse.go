package customerCourseController

import (
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/service/customerCourseService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerCourseController struct {
	CustomerCourseService customerCourseService.CustomerCourseService
}

// GetHistoryCourseByCustomerID is a function to get history course by customer id
func (ccc *CustomerCourseController) GetHistoryCourseByCustomerID(c echo.Context) error {
	// Get user id from jwt
	user := helper.GetUser(c)

	// Call service to get history course by customer id
	customerCourses, err := ccc.CustomerCourseService.GetHistoryCourseByCustomerID(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get history course",
			"error":   err.Error(),
		})
	}

	// return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success get history course",
		"courses": customerCourses,
	})
}

func (ccc *CustomerCourseController) TakeCourse(c echo.Context) error {
	var customerCourse dto.CustomerCourseTransaction
	// get course id from url
	courseID := c.Param("courseId")
	customerCourse.CourseID = courseID

	// validate request body
	if err := c.Validate(customerCourse); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Get customer id from jwt
	customer := helper.GetUser(c)
	customerCourse.CustomerID = customer.ID

	// call service to tak course
	err := ccc.CustomerCourseService.TakeCourse(customerCourse)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail take course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail take course",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success take course",
	})
}
