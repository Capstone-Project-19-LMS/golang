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

// DeleteCustomerCourse is a function to delete customer course
func (ccc *CustomerCourseController) DeleteCustomerCourse(c echo.Context) error {
	// get customer course id from url
	courseID := c.Param("courseId")

	// get customer id from jwt
	customer := helper.GetUser(c)

	// call service to delete customer course
	err := ccc.CustomerCourseService.DeleteCustomerCourse(courseID, customer.ID)
	if err != nil {	
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail delete customer course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete customer course",
			"error":   err.Error(),
		})
	}

	// return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete customer course",
	})
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

// GetCourseEnrollByID is a function to get course by id
func (ccc *CustomerCourseController) GetCustomerCourseEnrollByID(c echo.Context) error {
	// get id from url param
	id := c.Param("id")

	// get course by id from service
	customerEnroll, err := ccc.CustomerCourseService.GetCustomerCourseEnrollByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get customer course",
			"error":   err.Error(),
		})
	}

	// return response success
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success get customer course",
		"customer_enroll": customerEnroll,
	})
}

func (ccc *CustomerCourseController) TakeCourse(c echo.Context) error {
	var customerCourse dto.CustomerCourseTransaction
	// get course id from url
	courseID := c.Param("courseId")
	customerCourse.CourseID = courseID

	// Get customer id from jwt
	customer := helper.GetUser(c)
	customerCourse.CustomerID = customer.ID

	customerCourse.Status = true

	// validate request body
	if err := c.Validate(customerCourse); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}
	
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

func (ccc *CustomerCourseController) UpdateEnrollmentStatus(c echo.Context) error {
	var customerCourse dto.CustomerCourseTransaction
	// get course id from url
	id := c.Param("id")
	customerCourse.ID = id

	// get bind data from request body
	if err := c.Bind(&customerCourse); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "fail bind request body",
			"error":   err.Error(),
		})
	}

	// Get customer id from jwt
	instructor := helper.GetUser(c)

	// call service to update enrollment status
	err := ccc.CustomerCourseService.UpdateEnrollmentStatus(customerCourse, instructor.ID)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail update enrollment status",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail update enrollment status",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update enrollment status",
	})
}