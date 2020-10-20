package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ProductID uint `json:"product_id"`
	Title string `json:"title"`
	UserName string `json:"user_name"`
	Rating int `json:"rating"`
	Comments string `json:"comments"`
}
