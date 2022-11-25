package dto

type Instructor struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"profile_image"`
}

type InstructorResponseGet struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"profile_image"`
	Role         string `json:"role"`
}

type InstructorRegister struct {
	ID           string `json:"id"`
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	ProfileImage string `json:"profile_image"`
}

type InstructorLogin struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type InstructorResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
