# Golang Product & Reviews REST API Example
This repo is a RESTful API examples for Product & Reviews application with **Golang, Gorm and gorilla/mux**.

## Installation & Run
```bash
# Download this github repo
go get https://github.com/sudheerj/go-products-rest-demo

cd go-products-rest-demo
go build
./go-products-rest-demo
```

The API endpoints are accessible through `http://127.0.0.1:8081`

## Structure
```
├── server
│   ├── app.go
│   ├── handler          //API core handlers and response functions
│   │   ├── common.go    
│   │   ├── product.go   
│   │   └── review.go     
│── model
│       └── product.go     // Models for the application
│       └── review.go 
│── storage
│       └── product.go     // Data access layer for models
│       └── review.go 
│       └── storage.go
├── configs
│   └── config.go        // Configuration
└── main.go
```

## API

#### /products
* `GET` : Get all products
* `POST` : Create a new product

#### /products/:name
* `GET` : Get a product
* `PUT` : Update a product
* `DELETE` : Delete a product

#### /products/:name/reviews
* `GET` : Get all reviews of a product
* `POST` : Create a new review for a product

### /products/:name/reviews/:id
* `GET` : Get a review of a product
* `PUT` : Update a review of a product
* `DELETE` : Delete a review of a product