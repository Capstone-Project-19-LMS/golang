package quizcontroller

import (
	"golang/constant/constantError"
	"golang/models/dto"
	quizservice "golang/service/quizService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type QuizController struct {
	QuizService quizservice.QuizService
}

func (qc *QuizController) CreateQuiz(c echo.Context) error {
	var quiz dto.QuizTransaction
	// Binding request body to struct
	err := c.Bind(&quiz)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}

	// Validate request body
	if err = c.Validate(quiz); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "There is an empty field",
			"error":   err.Error(),
		})
	}

	// Call service to create quiz
	err = qc.QuizService.CreateQuiz(quiz)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail create quiz",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success create quiz",
	})
}

func (qc *QuizController) TakeQuiz(c echo.Context) error {
	// Get id from url
	var input dto.TakeQuizTransaction

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail bind data",
			"error":   err.Error(),
		})
	}
	// Call service to get module by id
	quiz, err := qc.QuizService.TakeQuiz(input)

	if err != nil {
		if _, ok := constantError.ErrorCode[err.Error()]; ok {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "fail get quiz by course id",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "fail get quiz by course id",
			"error":   err.Error(),
		})
	}

	// Return response if success
	return c.JSON(http.StatusOK, echo.Map{
		"message": "success get quiz by course id",
		"modules": quiz,
	})
}
