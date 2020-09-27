package storage

import (
	"fmt"
	"gorm.io/driver/mysql"
	"github.com/sudheerj/go-rest.git/configs"
	"github.com/sudheerj/go-rest.git/model"
	"gorm.io/gorm"
	"log"
)

type store interface {
	GetProducts(limit, offset int) (*[]model.Product, error)
	GetProduct(id int) (*model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(product *model.Product) error
	DeleteProduct(id int) error
}

type storeImpl struct {
	db *gorm.DB
}

func InitializeDB(dbConfig *configs.Config) {
	connectionString := fmt.Sprintf("%s:%s@/%s?charset=%s", dbConfig.DB.Username, dbConfig.DB.Password, dbConfig.DB.Database, dbConfig.DB.Charset)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	db = dbMigrate(db)
	initStore(&storeImpl{db: db})
	if err != nil {
		log.Fatal("could not connect storage", err)
	}

}

// dbMigrate will create and migrate the tables, and then make necessary relationships
func dbMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&model.Product{}, &model.Review{})
	// db.Model(&model.Review{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	return db
}

var DBStore store

func initStore(store store) {
	DBStore = store
}