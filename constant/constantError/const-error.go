package constantError

import "gorm.io/gorm"

const (
	// Error
	// ErrorNotAuthorized is error message when user not authorized
	ErrorNotAuthorized = "you are not authorized"
	// ErrorEmailOrPasswordNotMatch is error message when email or password not match
	ErrorEmailOrPasswordNotMatch = "email or password not match"
	// ErrorCategoryNotFound is error message when category not found
	ErrorCategoryNotFound = "category not found"
	// ErrorCourseCapacity is error message when course capacity is full
	ErrorCourseCapacity = "course capacity is full"
	// ErrorCustomerAlreadyTakeCourse is error message when customer already take course
	ErrorCustomerAlreadyTakeCourse     = "customer already take course"
	ErrorCustomerAlreadyFavoriteCourse = "customer already favorite the course"
	ErrorCustomerAlreadyRatingCourse   = "customer already review the course"
	ErrorCustomerNotFavoriteCourse     = "the customer is not favorite the course"
	ErrorCustomerNotRatingCourse       = "the customer is not review the course yet"
	ErrorCustomerNotEnrolled           = "the customer is not enrolled in the course"
	ErrorCourseNotFound                = "the course is not found"
	ErrorCapacityLowerThanZero         = "capacity lower than zero"
	ErrorCustomerNotFinishedCourse     = "the customer is not finished the course"
	ErrorDuplicateAssignmentCustomer   = "there is duplicate data in the customer assignment"
	ErrorAssignmentNotFoud             = "assignment not found"
	ErrorNoActive                      = "email not verifikasi"
)

var ErrorCode = map[string]int{
	gorm.ErrRecordNotFound.Error():               404,
	"category not found":                         404,
	"the customer is not favorite the course":    404,
	"the customer is not review the course yet":  404,
	"the course is not found":                    404,
	"you are not authorized":                     401,
	"email or password not match":                400,
	"course capacity is full":                    400,
	"customer already take course":               400,
	"customer already favorite the course":       400,
	"customer already review the course":         400,
	"the customer is not enrolled in the course": 400,
	"capacity lower than zero":                   400,
	"the customer is not finished the course":    400,
	"email not verifikasi":                       500,
}
