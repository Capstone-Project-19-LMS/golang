package quizrepository

import "golang/models/dto"

type QuizRepository interface {
	CreateQuiz(dto.QuizTransaction) error
	GetAllQuiz() ([]dto.Quiz, error)
	TakeQuiz(dto.TakeQuizTransaction) (dto.Quiz, error)
	DeleteQuiz(id string) error
}
