package assignmentservice

import (
	"golang/models/dto"
	assignmentmockrepository "golang/repository/assignmentRepository/assignmentMockRepository"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type suiteAssignment struct {
	suite.Suite
	assignmentService AssignmentService
	mock              *assignmentmockrepository.AssignmentMock
}

func (s *suiteAssignment) SetupTest() {
	mock := &assignmentmockrepository.AssignmentMock{}
	s.mock = mock
	NewAssignmentService := NewAssignmentService(s.mock)
	s.assignmentService = NewAssignmentService
}

func (s *suiteAssignment) TestGetAssignmentByID() {
	testCase := []struct {
		Name            string
		ParamID         string
		MockReturnBody  dto.Assignment
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    dto.Assignment
		ExpectedError   error
	}{
		{
			"success get assignment by id ",
			"abcde",
			dto.Assignment{
				ID:          "abcde",
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
				CustomerAssignments: []dto.CustomerAssignment{
					{
						ID:           "abcde",
						File:         "tes",
						Grade:        1,
						AssignmentID: "abcde",
						CustomerID:   "abcde",
					},
				},
			},
			nil,
			true,
			dto.Assignment{
				ID:          "abcde",
				Title:       "tes",
				Description: "tes",
				ModuleID:    "abcde",
				CustomerAssignments: []dto.CustomerAssignment{
					{
						ID:           "abcde",
						File:         "tes",
						Grade:        1,
						AssignmentID: "abcde",
						CustomerID:   "abcde",
					},
				},
			},
			nil,
		},
		{
			"failed get assignment by id",
			"abcde",

			dto.Assignment{},
			gorm.ErrRecordNotFound,
			false,
			dto.Assignment{},
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAssignmentByID", v.ParamID).Return(v.MockReturnBody, v.MockReturnError)
		s.T().Run(v.Name, func(t *testing.T) {
			assignment, err := s.assignmentService.GetAssignmentByID(v.ParamID)
			if v.HasReturnBody {
				s.NoError(err)
				s.Equal(v.ExpectedBody, assignment)
			} else {
				s.Error(err)
				s.EqualError(err, v.ExpectedError.Error())
				s.Equal(v.ExpectedBody, assignment)
			}
		})
		// remove mock
		mockCall.Unset()
	}
}

func TestSuiteAssignment(t *testing.T) {
	suite.Run(t, new(suiteAssignment))
}
