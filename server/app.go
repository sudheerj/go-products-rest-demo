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
	app.Router.HandleFunc("/product", handler.CreateProduct).Methods("POST")
	app.Router.HandleFunc("/product/{id:[0-9]+}", handler.GetProduct).Methods("GET")
	app.Router.HandleFunc("/product/{id:[0-9]+}", handler.UpdateProduct).Methods("PUT")
	app.Router.HandleFunc("/product/{id:[0-9]+}", handler.DeleteProduct).Methods("DELETE")
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}




