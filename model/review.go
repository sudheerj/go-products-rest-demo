package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ProductID int `json:"project_id"`
	Name string `json:"name"`
	Rating int `json:"rating"`
	Comments string `json:"comments"`
}
