package favoriteController

import (
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/service/favoriteService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FavoriteController struct {
	FavoriteService favoriteService.FavoriteService
}

func (fc *FavoriteController) AddFavorite(c echo.Context) error {
	var favorite dto.FavoriteTransaction
	// get course id from url
	courseID := c.Param("courseId")
	favorite.CourseID = courseID

	// validate request body
	if err := c.Validate(favorite); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Get customer id from jwt
	customer := helper.GetUser(c)
	favorite.CustomerID = customer.ID

	// call service to tak course
	err := fc.FavoriteService.AddFavorite(favorite)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail favorite course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail favorite course",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success favorite course",
	})
}

// DeleteFavorite is a function to delete customer course
func (fc *FavoriteController) DeleteFavorite(c echo.Context) error {
	// get customer course id from url
	courseID := c.Param("courseId")

	// get customer id from jwt
	customer := helper.GetUser(c)

	// call service to delete favorite course
	err := fc.FavoriteService.DeleteFavorite(courseID, customer.ID)
	if err != nil {	
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail delete favorite course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete favorite course",
			"error":   err.Error(),
		})
	}

	// return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete favorite course",
	})
}

// GetFavoriteCourseByCustomerID is a function to get favorite course by customer id
func (fc *FavoriteController) GetFavoriteCourseByCustomerID(c echo.Context) error {
	// Get user id from jwt
	user := helper.GetUser(c)

	// Call service to get favorite course by customer id
	favoriteCourses, err := fc.FavoriteService.GetFavoriteByCustomerID(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get favorite course",
			"error":   err.Error(),
		})
	}

	// return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success get favorite course",
		"courses": favoriteCourses,
	})
}