package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name                 string `json:"name" gorm:"notNull;size:255"`
	Email                string `json:"email" gorm:"notNull;unique;size:255"`
	Password             string `json:"password" gorm:"notNull"`
	ProfilePicture       string `json:"profile_picture" gorm:"size:255;default:null"`
}
