package customerassignmentcontroller

import (
	"golang/constant/constantError"
	"golang/models/dto"
	customerAssignmentService "golang/service/customerAssignmentService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerAssignmentController struct {
	CustomerAssignmentService customerAssignmentService.CustomerAssignmentService
}

// CreateCustomerAssignment is a function to create customerAssignment
func (cac *CustomerAssignmentController) CreateCustomerAssignment(c echo.Context) error {
	var customerAssignment dto.CustomerAssignmentTransaction
	// Binding request body to struct
	err := c.Bind(&customerAssignment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Validate request body
	if err = c.Validate(customerAssignment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Call service to create customerAssignment
	err = cac.CustomerAssignmentService.CreateCustomerAssignment(customerAssignment)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail create customer assignment",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create customer assignment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create customer assignment",
	})
}

// DeleteCustomerAssignment is a function to delete customer assignment
func (cac *CustomerAssignmentController) DeleteCustomerAssignment(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to delete customer assignment
	err := cac.CustomerAssignmentService.DeleteCustomerAssignment(id)
	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail delete customer assignment",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail delete customer assignment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success delete customer assignment",
	})
}

// GetAllCustomerAssignment is a function to get all customerAssignment
func (cac *CustomerAssignmentController) GetAllCustomerAssignment(c echo.Context) error {
	// Call service to get all customerAssignment
	customerAssignments, err := cac.CustomerAssignmentService.GetAllCustomerAssignment()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get all customer assignment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":             "success get all customer assignment",
		"customer_assignment": customerAssignments,
	})
}

// GetCustomerAssignmentByID is a function to get customerAssignment by id
func (cac *CustomerAssignmentController) GetCustomerAssignmentByID(c echo.Context) error {
	// Get id from url
	id := c.Param("id")

	// Call service to get customerAssignment by id
	customerAssignment, err := cac.CustomerAssignmentService.GetCustomerAssignmentByID(id)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail get customer assignment by id",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get customer assignment by id",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message":             "success get customer assignment by id",
		"customer_assignment": customerAssignment,
	})
}

// UpdateCustomerAssignment is a function to update customerAssignment
func (cac *CustomerAssignmentController) UpdateCustomerAssignment(c echo.Context) error {
	var customerAssignment dto.CustomerAssignmentTransaction
	// Binding request body to struct
	err := c.Bind(&customerAssignment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Get id from url
	id := c.Param("id")
	customerAssignment.ID = id

	// Call service to update customerAssignment
	err = cac.CustomerAssignmentService.UpdateCustomerAssignment(customerAssignment)
	if err != nil {
		if val, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(val, echo.Map{
				"message": "fail update customer assignment",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail update customer assignment",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update customer assignment",
	})
}
