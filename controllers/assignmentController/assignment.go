package assignmentcontroller

import (
	"golang/constant/constantError"
	"golang/models/dto"
	assignmentService "golang/service/assignmentService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AssignmentController struct {
	AssignmentService assignmentService.AssignmentService
}

// CreateAssignment is a function to create assignment
func (ac *AssignmentController) CreateAssignment(c echo.Context) error {
	var assignment dto.AssignmentTransaction
	// Binding request body to struct
	err := c.Bind(&assignment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Validate request body
	if err = c.Validate(assignment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Call service to create assignment
	err = ac.AssignmentService.CreateAssignment(assignment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create assignment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create assignment",
	})
}

// DeleteAssignment is a function to delete assignment
func (ac *AssignmentController) DeleteAssignment(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to delete assignment
	err := ac.AssignmentService.DeleteAssignment(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail delete assignment",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete assignment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete assignment",
	})
}

// GetAllAssignment is a function to get all assignment
func (ac *AssignmentController) GetAllAssignment(c echo.Context) error {
	// Call service to get all assignment
	assignments, err := ac.AssignmentService.GetAllAssignment()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get all assignment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":     "success get all assignment",
		"assignments": assignments,
	})
}

// GetAssignmentByID is a function to get assignment by id
func (ac *AssignmentController) GetAssignmentByID(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to get assignment by id
	assignment, err := ac.AssignmentService.GetAssignmentByID(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "failed get assignment by id",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "failed get assignment by id",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "success get assignment by id",
		"assignment": assignment,
	})
}

func (ac *AssignmentController) GetAssignmentByCourse(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to get assignment by id
	assignment, err := ac.AssignmentService.GetAssignmentByCourse(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "failed get assignment by course",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "failed get assignment by course",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":    "success get assignment by id",
		"assignment": assignment,
	})
}

// UpdateAssignment is a function to update assignment
func (ac *AssignmentController) UpdateAssignment(c echo.Context) error {
	var assignment dto.AssignmentTransaction
	// Binding request body to struct
	err := c.Bind(&assignment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Get id from url
	id := c.Param("id")
	assignment.ID = id

	// Call service to update assignment
	err = ac.AssignmentService.UpdateAssignment(assignment)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail update assignment",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail update assignment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update assignment",
	})
}
