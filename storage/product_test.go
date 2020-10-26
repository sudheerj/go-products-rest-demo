package storage

import (
	"github.com/sudheerj/go-rest.git/configs"
	"go/types"
	"strconv"
	"testing"

	"github.com/sudheerj/go-rest.git/model"
)

func getMockProduct(name string) model.Product {
	product := model.Product{
		Name: name,
		ModelType: "Audi A6",
		Price: 1000000,
		Color: "Red",
		Reviews: [] model.Review {
			model.Review{
				ProductID: 1,
				Title: "Good one",
				UserName: "john",
				Rating: 8,
				Comments: "Flexible and Comfortable",
			},
		},
	}
	return product
}

func createMockProduct(name string, t *testing.T)  {
	p := getMockProduct(name)
	err := DBStore.CreateProduct(&p)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestCreateProduct(t *testing.T) {
	p1 := getMockProduct("p1");
	type args struct {
		p *model.Product
	}
	tests := []struct {
		name string
		args args
		want types.Nil
		wantErr bool
	} {
		 {
		 	name: "product created successfully",
		 	args: args {
		 		p: &p1,
			},
			wantErr: false,
		 },
	}
	for _, testcase := range tests{
		t.Run(testcase.name, func(t *testing.T) {
			err := DBStore.CreateProduct(testcase.args.p)
			if (err != nil) != testcase.wantErr {
				t.Errorf("CreateProduct error=%v, wantErr %v", err, testcase.wantErr)
				return
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	for i:=0; i<3; i++ {
		productName := "p" + strconv.Itoa(i)
		createMockProduct(productName, t)
	}

	type args struct {
		name string
	}

	tests := []struct {
		name string
		args args
		wantErr bool
	} {
		{
			name: "product deleted successfully",
			args: args {
				name: "p1",
			},
			wantErr: false,
		},
		{
			name: "product is failed to delete",
			args: args {
				name: "p2",
			},
			wantErr: true,
		},
	}
	for _, testcase := range tests{
		t.Run(testcase.name, func(t *testing.T) {
			err := DBStore.DeleteProduct(testcase.args.name)
			if (err != nil) != testcase.wantErr {
				t.Errorf("DeleteProduct error=%v, wantErr %v", err, testcase.wantErr)
				return
			}
		})
	}

	deleteAllProducts()
}

func TestGetAllProducts(t *testing.T) {
	for i:=0; i<3; i++ {
		productName := "p" + strconv.Itoa(i)
		createMockProduct(productName, t)
	}

	type args struct {
		offset int
		limit int
	}

	tests := []struct {
		name string
		args args
		productsCount int
		wantErr bool
	} {
		{
			name: "Get default number of products",
			args: args {
				offset: 0,
			},
			productsCount: 10,
			wantErr: false,
		},
		{
			name: "Products with offset 0 and limit 4",
			args: args {
				offset: 0,
				limit: 4,
			},
			productsCount: 4,
			wantErr: false,
		},
		{
			name: "Products with offset 4 and limit 8",
			args: args {
				offset: 4,
				limit: 8,
			},
			productsCount: 4,
			wantErr: false,
		},
		{
			name: "Products with offset 10 and limit 5",
			args: args {
				offset: 10,
				limit: 5,
			},
			productsCount: 0,
			wantErr: false,
		},
	}
	for _, testcase := range tests{
		t.Run(testcase.name, func(t *testing.T) {
			got, err := DBStore.GetProducts(testcase.args.offset, testcase.args.limit)
			if (err != nil) != testcase.wantErr {
				t.Errorf("DeleteProduct error=%v, wantErr %v", err, testcase.wantErr)
				return
			}
			if testcase.productsCount != len(*got) {
				t.Errorf("Get Products error=%v, wantErr %v", len(*got), testcase.productsCount)
				return
			}
		})
	}

	deleteAllProducts()
}

func TestGetProduct(t *testing.T) {
	createMockProduct("p1", t)

	type args struct {
		name string
	}

	tests := []struct {
		name string
		args args
		wantName string
		wantErr bool
	} {
		{
			name: "Get product successfully",
			args: args {
				name: "p1",
			},
			wantName: "p1" ,
			wantErr: false,
		},
		{
			name: "Failed to get product",
			args: args {
				name: "p2",
			},
			wantName: "p3" ,
			wantErr: true,
		},
	}
	for _, testcase := range tests{
		t.Run(testcase.name, func(t *testing.T) {
			got, err := DBStore.GetProduct(testcase.args.name)
			if (err != nil) != testcase.wantErr {
				t.Errorf("Get Product error=%v, wantErr %v", err, testcase.wantErr)
				return
			}
			if got != nil && got.Name != testcase.wantName {
				t.Errorf("Get Products error=%v, wantErr %v", got.Name, testcase.wantName)
				return
			}
		})
	}

	deleteAllProducts()
}

func deleteAllProducts() {
	configs.DBConn.Unscoped().Exec("DELETE FROM products")
}


