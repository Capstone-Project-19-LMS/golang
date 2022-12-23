package quizrepository

import "golang/models/dto"

type QuizRepository interface {
	CreateQuiz(dto.QuizTransaction) error
	TakeQuiz(dto.TakeQuizTransaction) (dto.Quiz, error)
}
