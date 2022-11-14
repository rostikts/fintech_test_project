package service

import (
	"github.com/rostikts/fintech_test_project/config"
	"github.com/rostikts/fintech_test_project/db/models"
	"github.com/rostikts/fintech_test_project/internal/transaction/repository"
	"github.com/rostikts/fintech_test_project/test_utils"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
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

func TestGetDocumentsByMultipleFilters(t *testing.T) {
	data := models.Transaction{
		ID:              uint(rand.Intn(123456)),
		RequestID:       2,
		TerminalID:      3,
		PartnerObjectID: 44,
		Payment: models.Payment{
			Type:      "unique maybe payment type",
			Number:    "num",
			Narrative: "the sql injections fixed, i guess",
		},
		Service: models.Service{
			ID:   uint(rand.Intn(123456)),
			Name: "Test service",
		},
		Payee: models.Payee{
			ID:          uint(rand.Intn(123456)),
			Name:        "test",
			BankMfo:     123124,
			BankAccount: "acc",
		},
		AmountTotal:        20,
		AmountOriginal:     2,
		CommissionPS:       1.3,
		CommissionClient:   2,
		CommissionProvider: 4.1,
		DateInput:          time.Date(1999, 12, 12, 0, 0, 0, 0, time.Local),
		DatePost:           time.Date(1999, 12, 29, 0, 0, 0, 0, time.Local),
		Status:             "unique status for assert",
	}
	rawFilters := map[string]string{
		"terminal_id":       strconv.Itoa(int(data.TerminalID)),
		"status":            data.Status,
		"transaction_id":    strconv.Itoa(int(data.ID)),
		"payment_type":      data.Payment.Type,
		"from":              data.DateInput.Format("2006-01-02"),
		"to":                data.DatePost.Format("2006-01-02"),
		"payment_narrative": "fixed",
	}

	err := service.SaveTransaction(data)
	if err != nil {
		t.Fatalf("the transaction for test is not created\n%s", err.Error())
	}

	res, err := service.GetTransactions(rawFilters)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(res) != 1 {
		t.Error(res)
		t.Fatal("the incorrect number of elements is received")
	}
	t.Log(len(res))

	// match variable args (serial ids, dates)
	res[0].DateInput = data.DateInput
	res[0].DatePost = data.DatePost
	res[0].Payment.ID = 0

	if !reflect.DeepEqual(res[0], data) {
		t.Errorf("Returned incorrect element\n expected:%v\ngot: %v", data, res[0])
	}

}
