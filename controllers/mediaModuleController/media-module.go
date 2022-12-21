package mediamodulecontroller

import (
	"golang/constant/constantError"
	"golang/models/dto"
	mediamoduleservice "golang/service/mediaModuleService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MediaModuleController struct {
	MediaModuleService mediamoduleservice.MediaModuleService
}

// CreateMediaModule is a function to create module
func (mmc *MediaModuleController) CreateMediaModule(c echo.Context) error {
	var mediaModule dto.MediaModuleTransaction
	// Binding request body to struct
	err := c.Bind(&mediaModule)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Validate request body
	if err = c.Validate(mediaModule); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Call service to create module
	err = mmc.MediaModuleService.CreateMediaModule(mediaModule)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create media module",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create media module",
	})
}

// DeleteMediaModule is a function to delete media module
func (mmc *MediaModuleController) DeleteMediaModule(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to delete media module
	err := mmc.MediaModuleService.DeleteMediaModule(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail delete media module",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete media module",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete media module",
	})
}

// GetAllMediaModule is a function to get all module
func (mmc *MediaModuleController) GetAllMediaModule(c echo.Context) error {
	// Call service to get all module
	mediaModules, err := mmc.MediaModuleService.GetAllMediaModule()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get all media module",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":       "success get all media module",
		"media_modules": mediaModules,
	})
}

// GetMediaModuleByID is a function to get module by id
func (mmc *MediaModuleController) GetMediaModuleByID(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to get module by id
	mediaModule, err := mmc.MediaModuleService.GetMediaModuleByID(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail get media module by id",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get media module by id",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":      "success get media module by id",
		"media_module": mediaModule,
	})
}

// UpdateMediaModule is a function to update module
func (mmc *MediaModuleController) UpdateMediaModule(c echo.Context) error {
	var mediaModule dto.MediaModuleTransaction
	// Binding request body to struct
	err := c.Bind(&mediaModule)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Get id from url
	id := c.Param("id")
	mediaModule.ID = id

	// Call service to update module
	err = mmc.MediaModuleService.UpdateMediaModule(mediaModule)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail update media module",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail update media module",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update media module",
	})
}
