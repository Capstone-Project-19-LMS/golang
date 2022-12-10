package dto

type Costumer struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"profile_image" gorm:"size:255;default:null"`
	IsActive     bool   `json:"is_active"`
}

type CostumerResponseGet struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"profile_image" gorm:"size:255;default:null"`
	IsActive     bool   `json:"is_active"`
}

type CostumerRegister struct {
	ID             string `json:"id"`
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required"`
	ProfileImage   string `json:"profile_image" gorm:"size:255;default:null"`
	CustomerCodeID string `json:"customer_code_id"`
}

type CostumerLogin struct {
	ID       string `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CustomerVerif struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type CostumerResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
