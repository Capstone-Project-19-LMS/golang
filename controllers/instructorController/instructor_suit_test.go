package instructorcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"golang/helper"
	"golang/models/dto"
	instructormockservice "golang/service/instructorService/instructorMockService"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type suiteInstructor struct {
	suite.Suite
	instructorController *InstructorController
	mock                 *instructormockservice.InstructorMock
}

func (s *suiteInstructor) SetupTest() {
	mock := &instructormockservice.InstructorMock{}
	s.mock = mock
	s.instructorController = &InstructorController{
		InstructorService: s.mock,
	}
}

func (s *suiteInstructor) TestCreateInstructor() {
	testCase := []struct {
		Name               string
		Method             string
		Body               dto.InstructorRegister
		MockReturnError    error
		ExpectedStatusCode int
		ExpectedMesaage    string
	}{
		{
			"success create instructor",
			"POST",
			dto.InstructorRegister{
				ID:           "abcde",
				Name:         "tes",
				Email:        "tes@gmail.com",
				Password:     "tes123",
				ProfileImage: "tes",
			},
			nil,
			http.StatusOK,
			"success create instructor",
		},
		{
			"fail bind data",
			"POST",
			dto.InstructorRegister{
				Name:         "tes",
				Email:        "tes@gmail.com",
				Password:     "tes123",
				ProfileImage: "tes",
			},
			nil,
			http.StatusInternalServerError,
			"fail bind data",
		},
		{
			"There is an empty field",
			"POST",
			dto.InstructorRegister{
				Name:  "tes",
				Email: "tes@gmail.com",
			},
			nil,
			http.StatusBadRequest,
			"There is an empty field",
		},
		{
			"fail create instructor",
			"POST",
			dto.InstructorRegister{
				Name:         "tes",
				Email:        "tes@gmail.com",
				Password:     "tes123",
				ProfileImage: "tes",
			},

			errors.New("fail create instructor"),
			http.StatusInternalServerError,
			"fail create instructor",
		},
	}
	for i, v := range testCase {
		mockCall := s.mock.On("CreateInstructor", v.Body).Return(v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			res, _ := json.Marshal(v.Body)
			// Create request
			r := httptest.NewRequest(v.Method, "/instructor/", bytes.NewBuffer(res))
			if i != 1 {
				r.Header.Set("Content-Type", "application/json")
			}
			// Create response recorder
			w := httptest.NewRecorder()

			// handler echo
			e := echo.New()
			e.Validator = &helper.CustomValidator{
				Validator: validator.New(),
			}
			ctx := e.NewContext(r, w)
			ctx.SetPath("/instructor/create")

			err := s.instructorController.Register(ctx)
			s.NoError(err)
			s.Equal(v.ExpectedStatusCode, w.Code)

			var resp map[string]interface{}
			err = json.NewDecoder(w.Result().Body).Decode(&resp)
			s.NoError(err)

			s.Equal(v.ExpectedMesaage, resp["message"])
		})
		// remove mock
		mockCall.Unset()
	}
}

func TestSuiteInstructor(t *testing.T) {
	suite.Run(t, new(suiteInstructor))
}
