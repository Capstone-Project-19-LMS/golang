package assignmentservice

import (
	"golang/helper"
	"golang/models/dto"
	assignmentRepository "golang/repository/assignmentRepository"
)

type AssignmentService interface {
	CreateAssignment(dto.AssignmentTransaction) error
	DeleteAssignment(id string) error
	GetAllAssignment() ([]dto.Assignment, error)
	GetAssignmentByID(id string) (dto.Assignment, error)
	UpdateAssignment(dto.AssignmentTransaction) error
	GetAssignmentByCourse(id string) ([]dto.AssignmentCourse, error)
}

type assignmentService struct {
	assignmentRepo assignmentRepository.AssignmentRepository
}

// CreateModule implements ModuleService
func (as *assignmentService) CreateAssignment(assignment dto.AssignmentTransaction) error {
	id := helper.GenerateUUID()
	assignment.ID = id
	err := as.assignmentRepo.CreateAssignment(assignment)
	if err != nil {
		return err
	}
	return nil
}

// DeleteModule implements ModuleService
func (as *assignmentService) DeleteAssignment(id string) error {
	// call repository to delete account
	err := as.assignmentRepo.DeleteAssignment(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAllModule implements ModuleService
func (as *assignmentService) GetAllAssignment() ([]dto.Assignment, error) {
	assignments, err := as.assignmentRepo.GetAllAssignment()
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

// GetModuleByID implements ModuleService
func (as *assignmentService) GetAssignmentByID(id string) (dto.Assignment, error) {
	assignment, err := as.assignmentRepo.GetAssignmentByID(id)
	if err != nil {
		return dto.Assignment{}, err
	}
	return assignment, nil
}
func (as *assignmentService) GetAssignmentByCourse(id string) ([]dto.AssignmentCourse, error) {
	assignments, err := as.assignmentRepo.GetAssignmentByCourse(id)
	if err != nil {
		return []dto.AssignmentCourse{}, err
	}
	return assignments, nil
}

// UpdateModule implements ModuleService
func (as *assignmentService) UpdateAssignment(assignment dto.AssignmentTransaction) error {
	// call repository to update Module
	err := as.assignmentRepo.UpdateAssignment(assignment)
	if err != nil {
		return err
	}
	return nil
}

func NewAssignmentService(assignmentRepo assignmentRepository.AssignmentRepository) AssignmentService {
	return &assignmentService{
		assignmentRepo: assignmentRepo,
	}
}
