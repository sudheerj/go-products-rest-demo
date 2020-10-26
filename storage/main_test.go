package storage

import (
	"fmt"
	"github.com/sudheerj/go-rest.git/configs"
	"github.com/sudheerj/go-rest.git/model"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	dbConfig := configs.GetConfig()
	InitializeDB(dbConfig)

	exitVal := m.Run()

	if err := configs.DBConn.Migrator().DropTable(&model.Product{}, &model.Review{}); err != nil {
		fmt.Println("failed to cleanup")
	}

	os.Exit(exitVal)
}
