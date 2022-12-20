package categoryController

import (
	"golang/constant/constantError"
	"golang/helper"
	"golang/models/dto"
	"golang/service/categoryService"
	"net/http"

	"github.com/jinzhu/copier"
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

// DeleteCategory is a function to delete category
func (cc *CategoryController) DeleteCategory(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to delete category
	err := cc.CategoryService.DeleteCategory(id)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail delete category",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete category",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete category",
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

	// Get user id from jwt
	user := helper.GetUser(c)
	
	// Call service to get category by id
	getCategory, err := cc.CategoryService.GetCategoryByID(id, user)
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

	if user.Role == "instructor" {
		var categoryInstructor dto.GetCategoryInstructor
		_ = copier.Copy(&categoryInstructor, &getCategory)
		// Return response if success
		return c.JSON(http.StatusOK, echo.Map{
			"message":   "success get category by id",
			"category": categoryInstructor,
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get category by id",
		"category": getCategory,
	})
}

// UpdateCategory is a function to update category
func (cc *CategoryController) UpdateCategory(c echo.Context) error {
	var category dto.CategoryTransaction
	// Binding request body to struct
	err := c.Bind(&category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Get id from url
	id := c.Param("id")
	category.ID = id

	// Call service to update category
	err = cc.CategoryService.UpdateCategory(category)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail update category",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail update category",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update category",
	})
}