package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Item      string `json:"item" gorm:"text;not null"`
	Completed bool   `json:"completed" gorm:"bool;default:false"`
}
