package storage

import (
	"errors"
	"github.com/sudheerj/go-rest.git/model"
	"gorm.io/gorm"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func (storeImpl *storeImpl) GetProducts(offset, limit int) (*[]model.Product, error) {
	log.WithFields(
		log.Fields{
			"offset": offset,
			"limit": limit,
		},
	).Info("Get products")
	products := make([]model.Product, 0)
	if err := storeImpl.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return &products, nil
}

func (storeImpl *storeImpl) GetProduct(name string) (*model.Product, error) {
	log.WithFields(
		log.Fields{
			"ProductName": name,
		},
	).Info("Get product")

    product := model.Product{}
    if err := storeImpl.db.First(&product, model.Product{Name: name}).Error; err != nil {
    	return nil, err
	}
    return &product, nil
}

func (storeImpl *storeImpl) CreateProduct(product *model.Product) error {
	log.Println("Create new product")

	if product == nil {
		return errors.New("product information is mandatory")
	}
	if err := storeImpl.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) UpdateProduct(product *model.Product) error {
	log.WithFields(
		log.Fields{
			"ID": product.Name,
		},
	).Info("Update product")

	if product == nil {
		return errors.New("no product found with given product name")
	}
	if err := storeImpl.db.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func (storeImpl *storeImpl) DeleteProduct(name string) error {
	log.WithFields(
		log.Fields{
			"ProductName": name,
		},
	).Info("Delete product")

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
	log.Println("Get all products")
	products := make([]model.Product, 0)
	db.Find(&products)
	return products
}