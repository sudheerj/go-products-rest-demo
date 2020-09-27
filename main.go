package main

import (
	"github.com/sudheerj/go-rest.git/configs"
	"github.com/sudheerj/go-rest.git/server"
	"github.com/sudheerj/go-rest.git/storage"
)

func main() {
	dbConfig := configs.GetConfig()

	app := &server.App{}
	storage.InitializeDB(dbConfig)
	app.InitializeRoutes()
	app.Run(":8081")
}
