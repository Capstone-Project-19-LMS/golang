package quizrepository

import (
	"errors"
	"fmt"
	"golang/constant/constantError"
	"golang/models/dto"
	"golang/models/model"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type quizRepository struct {
	db *gorm.DB
}

func (ctr *quizRepository) CreateQuiz(input dto.QuizTransaction) error {
	var quizModel model.Quiz
	err := copier.Copy(&quizModel, &input)
	if err != nil {
		return err
	}

	err = ctr.db.Create(&quizModel).Error
	if err != nil {
		return err
	}
	return nil

}

func (ctr *quizRepository) GetAllQuiz() ([]dto.Quiz, error) {
	var quizModel []model.Quiz
	var quiz []dto.Quiz
	err := ctr.db.Find(&quizModel).Error
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&quiz, quizModel)
	if err != nil {
		return nil, err
	}
	return quiz, nil

}

func (ctr *quizRepository) TakeQuiz(input dto.TakeQuizTransaction) (dto.Quiz, error) {
	var quizModel model.Quiz
	var customerCourse model.CustomerCourse
	var quiz dto.Quiz
	ctr.db.Where("customer_id=?", input.CustomerID).Where("course_id=?", input.CourseID).Find(&customerCourse)
	fmt.Println(customerCourse)
	fmt.Println(input.CustomerID)
	fmt.Println(input.CourseID)
	if !customerCourse.Status {

		return dto.Quiz{}, errors.New(constantError.ErrorCustomerNotEnrolled)
	}
	fmt.Println(customerCourse.Status)
	ctr.db.Where("course_id=?", input.CourseID).Find(&quizModel)
	err := copier.Copy(&quiz, quizModel)
	if err != nil {
		return dto.Quiz{}, err
	}
	return quiz, nil
}

func (ctr *quizRepository) DeleteQuiz(id string) error {

	err := ctr.db.Where("id = ?", id).Unscoped().Delete(&model.Quiz{})
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func NewQuizRepository(db *gorm.DB) QuizRepository {
	return &quizRepository{
		db: db,
	}
}
