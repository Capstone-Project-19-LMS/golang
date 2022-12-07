package ratingController

import (
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/service/ratingService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RatingController struct {
	RatingService ratingService.RatingService
}

func (rc *RatingController) AddRating(c echo.Context) error {
	var rating dto.RatingTransaction
	// get course id from url
	courseID := c.Param("courseId")
	rating.CourseID = courseID

	// Binding request body to struct
	err := c.Bind(&rating)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// validate request body
	if err := c.Validate(rating); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Get customer id from jwt
	customer := helper.GetUser(c)
	rating.CustomerID = customer.ID

	// call service to review the course
	err = rc.RatingService.AddRating(rating)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail rating course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail rating course",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success rating course",
	})
}

// DeleteRating is a function to delete rating by customer
func (rc *RatingController) DeleteRating(c echo.Context) error {
	// get course id from url
	courseID := c.Param("courseId")

	// get customer id from jwt
	customer := helper.GetUser(c)

	// call service to delete rating course
	err := rc.RatingService.DeleteRating(courseID, customer.ID)
	if err != nil {	
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail delete rating course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete rating course",
			"error":   err.Error(),
		})
	}

	// return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete rating course",
	})
}