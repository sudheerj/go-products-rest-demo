package main

import (
	"os"
	log "github.com/sirupsen/logrus"

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
	log.Println("Server runs on 8081 port")
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}
