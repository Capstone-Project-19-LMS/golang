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
	ErrorCustomerAlreadyTakeCourse = "customer already take course"
)

var ErrorCode = map[string]int{
	gorm.ErrRecordNotFound.Error():   404,
	"category not found":             404,
	"you are not authorized":         401,
	"email or password not match": 400,
	"course capacity is full":        400,
	"customer already take course":   400,
	ErrorNoActive = "email not verifikasi"
)
