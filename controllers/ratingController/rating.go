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

// GetRating is a function to get a rating by customer
func (rc *RatingController) GetRatingByCourseIDCustomerID(c echo.Context) error {
	// get course id from url
	courseID := c.Param("courseId")

	// get customer id from jwt
	customer := helper.GetUser(c)

	// call service to get rating course
	rating, err := rc.RatingService.GetRatingByCourseIDCustomerID(courseID, customer.ID)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail get rating course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get rating course",
			"error":   err.Error(),
		})
	}

	// return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get rating course",
		"data":    rating,
	})
}

// GetRatingByCourseID is a function to get all rating by course id
func (rc *RatingController) GetRatingByCourseID(c echo.Context) error {
	// get course id from url
	courseID := c.Param("courseId")

	// get customer id from jwt
	instructor := helper.GetUser(c)

	// call service to get rating course
	rating, err := rc.RatingService.GetRatingByCourseID(courseID, instructor.ID)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail get rating of course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get rating of course",
			"error":   err.Error(),
		})
	}

	// return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get rating of course",
		"rating":    rating,
	})
}

// UpdateRating is a function to update rating by customer
func (rc *RatingController) UpdateRating(c echo.Context) error {
	// get rating id from url
	ratingID := c.Param("ratingId")

	var rating dto.RatingTransaction
	// Binding request body to struct
	err := c.Bind(&rating)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	rating.ID = ratingID

	// call service to update rating course
	err = rc.RatingService.UpdateRating(rating)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail update rating course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail update rating course",
			"error":   err.Error(),
		})
	}

	// return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update rating course",
	})
}