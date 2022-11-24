package categoryController

import (
	"golang/constant/constantError"
	"golang/models/dto"
	"golang/service/categoryService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryService categoryService.CategoryService
}

// CreateCategory is a function to create category
func (cc *CategoryController) CreateCategory(c echo.Context) error {
	var category dto.Category
	// Binding request body to struct
	err := c.Bind(&category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Validate request body
	if err = c.Validate(category); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Call service to create category
	err = cc.CategoryService.CreateCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create category",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create category",
	})
}

// DeleteCategory is a function to delete account
func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to delete account
	err := cc.CategoryService.DeleteCategory(id)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail delete account",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete account",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete account",
	})
}