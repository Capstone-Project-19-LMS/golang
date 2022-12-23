package courseController

import (
	middlewares "golang/app/middlewares/instructor"
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/service/courseService"
	"net/http"

	"github.com/jinzhu/copier"
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
	user := helper.GetUser(c)

	// Call service to create course
	err = cc.CourseService.CreateCourse(course, user)
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

// DeleteCourse is a function to delete course
func (cc *CourseController) DeleteCourse(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Get instructor id from jwt
	claims := middlewares.GetUserInstructor(c)
	instructorId := claims.ID

	// Call service to delete course
	err := cc.CourseService.DeleteCourse(id, instructorId)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail delete course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete course",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete course",
	})
}

// GetAllCourse is a function to get all course
func (cc *CourseController) GetAllCourse(c echo.Context) error {
	// Get user id from jwt
	user := helper.GetUser(c)

	// Call service to get all courses
	getCourses, err := cc.CourseService.GetAllCourse(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get all courses",
			"error":   err.Error(),
		})
	}

	if user.Role == "instructor" {
		var coursesInstructor []dto.GetCourseInstructor
		err = copier.Copy(&coursesInstructor, &getCourses)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail get all courses",
				"error":   err.Error(),
			})
		}
		var amountModule int
		var amountCustomer int
		for _, course := range coursesInstructor {
			amountModule += course.NumberOfModules
			amountCustomer += course.AmountCustomer
		}
		// Return response if success
		return c.JSON(http.StatusOK, echo.Map{
			"message":   "success get all courses",
			"courses": coursesInstructor,
			"amount_course":len(coursesInstructor),
			"amount_materi" : amountModule,
			"amount_customer_course" : amountCustomer,
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success get all courses",
		"courses": getCourses,
	})
}

// GetCourseByID is a function to get course by id
func (cc *CourseController) GetCourseByID(c echo.Context) error {
	// get id from url param
	id := c.Param("id")

	// Get user id from jwt
	user := helper.GetUser(c)

	// get course by id from service
	getCourses, err := cc.CourseService.GetCourseByID(id, user)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail get course by id",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get course by id",
			"error":   err.Error(),
		})
	}

	if user.Role == "instructor" {
		var courseInstructor dto.GetCourseInstructorByID
		err = copier.Copy(&courseInstructor, &getCourses)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail get all courses",
				"error":   err.Error(),
			})
		}
		// Return response if success
		return c.JSON(http.StatusOK, echo.Map{
			"message":   "success get course by id",
			"course": courseInstructor,
		})
	}

	// return response success
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success get course by id",
		"course": getCourses,
	})
}

// GetCourseEnrollByID is a function to get course by id
func (cc *CourseController) GetCourseEnrollByID(c echo.Context) error {
	// get id from url param
	id := c.Param("courseId")

	// Get user id from jwt
	user := helper.GetUser(c)

	// get course by id from service
	customerEnroll, err := cc.CourseService.GetCourseEnrollByID(id, user)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail get course with customer enrolled",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get course with customer enrolled",
			"error":   err.Error(),
		})
	}

	// return response success
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success get course with customer enrolled",
		"customer_enroll": customerEnroll,
	})
}

// UpdateCourse is a function to update course
func (cc *CourseController) UpdateCourse(c echo.Context) error {
	var course dto.CourseTransaction
	// Binding request body to struct
	err := c.Bind(&course)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// get id from url
	id := c.Param("id")
	course.ID = id

	// Get user id from jwt
	claims := middlewares.GetUserInstructor(c)
	course.InstructorID = claims.ID

	// Call service to update course
	err = cc.CourseService.UpdateCourse(course)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail update course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail update course",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update course",
	})
}