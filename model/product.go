package model

type Product struct {
	ID   int    `json:"id"`
	Name	string `json:"name"`
	ModelType string `json:"modelType"`
	Price  uint32    `json:"price"`
	Color string `json:"color"`
	Reviews    []Review `gorm:"ForeignKey:ProductID" json:"reviews"`
}


