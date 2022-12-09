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
		ParamUser       dto.User
		MockReturnBody  dto.Assignment
		MockReturnError error
		HasReturnBody   bool
		ExpectedBody    dto.GetAssignment
		ExpectedError   error
	}{
		{
			"success get assignment by id and customer",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.Assignment{
				ID:          "abcde",
				Title:       "tes",
				Description: "tes",
				ModuleID:    "tes",
				CustomerAssignments: []dto.CustomerAssignment{
					{
						ID:           "tes",
						File:         "tes",
						Grade:        1,
						AssignmentID: "tes",
						CustomerID:   "tes",
					},
				},
			},
			nil,
			true,
			dto.GetAssignment{
				ID:          "abcde",
				Title:       "tes",
				Description: "tes",
				ModuleID:    "tes",
				CustomerAssignments: []dto.CustomerAssignment{
					{
						ID:           "tes",
						File:         "tes",
						Grade:        1,
						AssignmentID: "tes",
						CustomerID:   "tes",
					},
				},
			},
			nil,
		},
		{
			"failed get assignment by id and customer",
			"abcde",
			dto.User{
				ID:   "abcde",
				Role: "customer",
			},
			dto.Assignment{},
			gorm.ErrRecordNotFound,
			false,
			dto.GetAssignment{},
			gorm.ErrRecordNotFound,
		},
	}
	for _, v := range testCase {
		mockCall := s.mock.On("GetAssignmentByID", v.ParamID, v.ParamUser).Return(v.MockReturnBody, v.MockReturnError)
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
