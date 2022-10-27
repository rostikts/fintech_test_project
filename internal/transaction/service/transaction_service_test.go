package service

import (
	"github.com/rostikts/fintech_test_project/config"
	"github.com/rostikts/fintech_test_project/internal/transaction/repository"
	"github.com/rostikts/fintech_test_project/test_utils"
	"os"
	"testing"
)

const url = `https://drive.google.com/u/0/uc?id=1IwZ3uUCHGpSL2OoQu4mtbw7Ew3ZamcGB&export=download`

var service transactionService

func TestMain(m *testing.M) {
	db := test_utils.NewTestDB(config.Config.DBConfig)
	repo := repository.NewTransactionRepository(db)
	service = transactionService{repo: repo}
	defer db.Close()
	os.Exit(m.Run())
}

func TestLoaderService_ParseDocument(t *testing.T) {
	success, failed, err := service.ParseDocument(url)
	if err != nil {
		t.Error(err)
	}
	if success != 100 || failed != 0 {
		t.Error("The incorrect number of documents are saved in db")
	}

}
