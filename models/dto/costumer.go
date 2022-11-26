package dto

type Costumer struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"profile_image" gorm:"size:255;default:null"`
}

type CostumerResponseGet struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"profile_image" gorm:"size:255;default:null"`
	Role         string `json:"role"`
}

type CostumerRegister struct {
	ID           string `json:"id"`
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
	ProfileImage string `json:"profile_image" gorm:"size:255;default:null"`
}

type CostumerLogin struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CostumerResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
