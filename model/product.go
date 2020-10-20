package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name	string `json:"name" gorm:"unique"`
	ModelType string `json:"modelType"`
	Price  uint32    `json:"price"`
	Color string `json:"color"`
	Reviews    []Review `gorm:"ForeignKey:ProductID" json:"reviews"`
}


