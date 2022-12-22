package assignmentrepository

import "golang/models/dto"

type AssignmentRepository interface {
	CreateAssignment(dto.AssignmentTransaction) error
	DeleteAssignment(id string) error
	GetAllAssignment() ([]dto.Assignment, error)
	GetAssignmentByID(id string) (dto.Assignment, error)
	UpdateAssignment(dto.AssignmentTransaction) error
	GetAssignmentByCourse(id string) ([]dto.AssignmentCourse, error)
}
