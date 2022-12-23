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
	copier.Copy(&quiz, quizModel)
	return quiz, nil
}

func NewQuizRepository(db *gorm.DB) QuizRepository {
	return &quizRepository{
		db: db,
	}
}
