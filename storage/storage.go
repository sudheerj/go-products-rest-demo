package storage

import (
	"fmt"
	"github.com/sudheerj/go-rest.git/configs"
	"github.com/sudheerj/go-rest.git/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type store interface {
	GetProducts(offset, limit int) (*[]model.Product, error)
	GetProduct(name string) (*model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(name string) error

	GetReviews(product *model.Product) (*[]model.Review, error)
	GetReview(id int) (*model.Review, error)
	CreateReview(newReview *model.Review) error
	UpdateReview(review *model.Review) error
	DeleteReview(review *model.Review) error
}

type storeImpl struct {
	db *gorm.DB
}

func InitializeDB(dbConfig *configs.Config) {
	connectionString := fmt.Sprintf("%s:%s@/%s?charset=%s", dbConfig.DB.Username, dbConfig.DB.Password, dbConfig.DB.Name, dbConfig.DB.Charset)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	db = dbMigrate(db)
	initStore(&storeImpl{db: db})
	if err != nil {
		log.Fatal("could not connect storage", err)
	}

}

// dbMigrate will create and migrate the tables, and then make necessary relationships
func dbMigrate(db *gorm.DB) *gorm.DB {
	if migrationError := db.AutoMigrate(&model.Product{}, &model.Review{}); migrationError != nil {
		log.Fatal(migrationError);
	}
	return db
}

var DBStore store

func initStore(store store) {
	DBStore = store
}