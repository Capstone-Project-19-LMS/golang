package quizservice

import (
	"golang/helper"
	"golang/models/dto"
	quizrepository "golang/repository/quizRepository"
)

type QuizService interface {
	CreateQuiz(dto.QuizTransaction) error
	TakeQuiz(dto.TakeQuizTransaction) (dto.Quiz, error)
}

type quizService struct {
	quizRepo quizrepository.QuizRepository
}

// CreateCustomerAssignment implements QuizService
func (cas *quizService) CreateQuiz(input dto.QuizTransaction) error {
	id := helper.GenerateUUID()
	input.ID = id
	err := cas.quizRepo.CreateQuiz(input)
	if err != nil {
		return err
	}
	return nil
}

func (cas *quizService) TakeQuiz(input dto.TakeQuizTransaction) (dto.Quiz, error) {
	quiz, err := cas.quizRepo.TakeQuiz(input)
	if err != nil {
		return dto.Quiz{}, err
	}
	return quiz, nil
}

func NewQuizService(quizRepo quizrepository.QuizRepository) QuizService {
	return &quizService{
		quizRepo: quizRepo,
	}
}
