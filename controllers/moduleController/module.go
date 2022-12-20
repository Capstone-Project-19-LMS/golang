package moduleController

import (
	"golang/constant/constantError"
	"golang/models/dto"
	"golang/models/model"
	moduleservice "golang/service/moduleService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ModuleController struct {
	ModuleService moduleservice.ModuleService
}

// CreateModule is a function to create module
func (mc *ModuleController) CreateModule(c echo.Context) error {
	var module dto.ModuleTransaction
	// Binding request body to struct
	err := c.Bind(&module)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Validate request body
	if err = c.Validate(module); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Call service to create module
	err = mc.ModuleService.CreateModule(module)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create module",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create module",
	})
}

// DeleteModule is a function to delete module
func (mc *ModuleController) DeleteModule(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to delete module
	err := mc.ModuleService.DeleteModule(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail delete module",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete module",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete module",
	})
}

// GetAllModule is a function to get all module
func (mc *ModuleController) GetAllModule(c echo.Context) error {
	// Call service to get all module
	modules, err := mc.ModuleService.GetAllModule()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get all module",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get all module",
		"modules": modules,
	})
}

// GetModuleByID is a function to get module by id
func (mc *ModuleController) GetModuleByID(c echo.Context) error {
	// Get id from url
	id := c.Param("id")
	input := new(model.CustomerCourse)
	err := c.Bind(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}
	// Call service to get module by id
	module, err := mc.ModuleService.GetModuleByID(id, input.CustomerID)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail get module by id",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get module by id",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get module by id",
		"module":  module,
	})
}

// GetModuleByCourseID is a function to get module by id
func (mc *ModuleController) GetModuleByCourseID(c echo.Context) error {
	// Get id from url
	input := new(model.CustomerCourse)

	err := c.Bind(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}
	// Call service to get module by id
	modules, err := mc.ModuleService.GetModuleByCourseID(input.CourseID, input.CustomerID)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail get module by course id",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get module by course id",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get module by id",
		"modules": modules,
	})
}

// UpdateModule is a function to update module
func (mc *ModuleController) UpdateModule(c echo.Context) error {
	var module dto.ModuleTransaction
	// Binding request body to struct
	err := c.Bind(&module)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Get id from url
	id := c.Param("id")
	module.ID = id

	// Call service to update module
	err = mc.ModuleService.UpdateModule(module)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail update module",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail update module",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update module",
	})
}
