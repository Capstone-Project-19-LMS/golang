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
	var category dto.CategoryTransaction
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

// GetAllCategory is a function to get all category
func (cc *CategoryController) GetAllCategory(c echo.Context) error {
	// Call service to get all category
	categories, err := cc.CategoryService.GetAllCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get all category",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "success get all category",
		"categories": categories,
	})
}

// GetCategoryByID is a function to get category by id
func (cc *CategoryController) GetCategoryByID(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to get category by id
	category, err := cc.CategoryService.GetCategoryByID(id)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail get category by id",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get category by id",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get category by id",
		"category": category,
	})
}