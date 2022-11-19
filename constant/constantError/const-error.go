package constantError

import "gorm.io/gorm"

const (
	// Error
	// ErrorNotAuthorized is error message when user not authorized
	ErrorNotAuthorized = "you are not authorized"
	// ErrorEmailOrPasswordNotMatch is error message when email or password not match
	ErrorEmailOrPasswordNotMatch = "email or password not match"
)

var ErrorCode = map[string]int{
	gorm.ErrRecordNotFound.Error():   404,
	"you are not authorized":         401,
	"email or password not match": 400,
}
