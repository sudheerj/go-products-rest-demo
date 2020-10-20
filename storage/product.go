package storage

import (
	"errors"
	"github.com/sudheerj/go-rest.git/model"
	"gorm.io/gorm"
	"net/http"
)

func (storeImpl *storeImpl) GetProducts(offset, limit int) (*[]model.Product, error) {
	products := make([]model.Product, 0)
	if err := storeImpl.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (storeImpl *storeImpl) GetProduct(name string) (*model.Product, error) {
    product := model.Product{}
    if err := storeImpl.db.First(&product, model.Product{Name: name}).Error; err != nil {
    	return nil, err
	}
    return &product, nil
}

func (storeImpl *storeImpl) CreateProduct(product *model.Product) error {
	if product == nil {
		return errors.New("product information is mandatory")
	}
	if err := storeImpl.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) UpdateProduct(product *model.Product) error {
	if product == nil {
		return errors.New("no product found with given product name")
	}
	if err := storeImpl.db.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) DeleteProduct(name string) error {
	product, err := storeImpl.GetProduct(name)
	if err != nil {
		return errors.New("no product found with given product name")
	}
	if err := storeImpl.db.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) []model.Product {
	products := make([]model.Product, 0)
	db.Find(&products)
	return products
}