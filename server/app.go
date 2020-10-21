package server

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sudheerj/go-rest.git/server/handler"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (app *App) InitializeRoutes() {
	app.Router = mux.NewRouter()

	app.Router.HandleFunc("/products", handler.GetProducts).Methods("GET")
	app.Router.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	app.Router.HandleFunc("/products/{name}", handler.GetProduct).Methods("GET")
	app.Router.HandleFunc("/products/{name}", handler.UpdateProduct).Methods("PUT")
	app.Router.HandleFunc("/products/{name}", handler.DeleteProduct).Methods("DELETE")

	app.Router.HandleFunc("/products/{name}/reviews", handler.GetReviews).Methods("GET")
	app.Router.HandleFunc("/products/{name}/reviews", handler.CreateReview).Methods("POST")
	app.Router.HandleFunc("/products/{name}/reviews/{id}", handler.GetReview).Methods("GET")
	app.Router.HandleFunc("/products/{name}/reviews/{id}", handler.UpdateReview).Methods("PUT")
	app.Router.HandleFunc("/products/{name}/reviews/{id}", handler.DeleteReview).Methods("DELETE")
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}




