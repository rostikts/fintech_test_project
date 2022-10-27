package repository

import (
	"github.com/rostikts/fintech_test_project/config"
	"github.com/rostikts/fintech_test_project/db/models"
	"github.com/rostikts/fintech_test_project/test_utils"
	"os"
	"reflect"
	"testing"
	"time"
)

var repo transactionRepository

func TestMain(m *testing.M) {
	db := test_utils.NewTestDB(config.Config.DBConfig)
	repo = transactionRepository{db: db}
	defer db.Close()
	os.Exit(m.Run())
}

func TestSaveTransaction(t *testing.T) {
	data := models.Transaction{
		RequestID:       2,
		TerminalID:      3,
		PartnerObjectID: 4,
		Payment: models.Payment{
			ID:        1,
			Type:      "type",
			Number:    "num",
			Narrative: "narrative",
		},
		Service: models.Service{
			ID:   1,
			Name: "Test service",
		},
		Payee: models.Payee{
			ID:          1,
			Name:        "test",
			BankMfo:     123124,
			BankAccount: "acc",
		},
		AmountTotal:        20,
		AmountOriginal:     2,
		CommissionPS:       1.3,
		CommissionClient:   2,
		CommissionProvider: 4.1,
		DateInput:          time.Now(),
		DatePost:           time.Now(),
		Status:             "not cool",
	}
	err := repo.SaveTransaction(data)
	if err != nil {
		t.Errorf("The transaction is not saved due to the error: %v", err)
	}
}

func TestTransactionRepositoryGetRecords(t *testing.T) {
	data := models.Transaction{
		RequestID:       2,
		TerminalID:      3,
		PartnerObjectID: 4,
		Payment: models.Payment{
			ID:        2,
			Type:      "type",
			Number:    "num",
			Narrative: "narrative",
		},
		Service: models.Service{
			ID:   2,
			Name: "Test service",
		},
		Payee: models.Payee{
			ID:          2,
			Name:        "test",
			BankMfo:     123124,
			BankAccount: "acc",
		},
		AmountTotal:        20,
		AmountOriginal:     2,
		CommissionPS:       1.3,
		CommissionClient:   2,
		CommissionProvider: 4.1,
		DateInput:          time.Now(),
		DatePost:           time.Now(),
		Status:             "unique status for assert",
	}
	if err := repo.SaveTransaction(data); err != nil {
		t.Fatalf("The transaction is not saved due to the error: %v", err)
	}
	res, err := repo.GetRecords()
	if err != nil {
		t.Fatalf("The list of transactions are not found due to the error: %v", err)
	}
	for _, v := range res {

		// exclude various fields
		v.ID = 0
		v.DateInput = data.DateInput
		v.DatePost = data.DatePost

		if reflect.DeepEqual(v, data) {
			return
		}
	}

	t.Errorf("Created transaction is not found in the result list")
}
