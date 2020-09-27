package storage

import (
	"errors"
	"github.com/sudheerj/go-rest.git/model"
	"gorm.io/gorm"
	"net/http"
)

func (storeImpl *storeImpl) GetProducts(limit, offset int) (*[]model.Product, error) {
	products := []model.Product{}
	if err := storeImpl.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (storeImpl *storeImpl) GetProduct(id int) (*model.Product, error) {
    product := model.Product{}
    if err := storeImpl.db.First(&product, model.Product{ID: id}).Error; err != nil {
    	return nil, err
	}
    return &product, nil
}

func (storeImpl *storeImpl) CreateProduct(product *model.Product) error {
	if product == nil {
		return errors.New("no product found with given product Id")
	}
	if err := storeImpl.db.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) UpdateProduct(product *model.Product) error {
	if product == nil {
		return errors.New("no product found with given product Id")
	}
	if err := storeImpl.db.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) DeleteProduct(id int) error {
	product, err := storeImpl.GetProduct(id)
	if err != nil {
		return errors.New("no product found with given product Id")
	}
	if err := storeImpl.db.Delete(&product).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) []model.Product {
	products := []model.Product{}
	db.Find(&products)
	return products
}