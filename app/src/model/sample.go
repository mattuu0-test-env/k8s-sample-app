package model

import "gorm.io/gorm"

type Sample struct {
	gorm.Model
	Message string `json:"message"`
}
