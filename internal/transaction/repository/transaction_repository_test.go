package repository

import (
	"fmt"
	"github.com/rostikts/fintech_test_project/config"
	"github.com/rostikts/fintech_test_project/db/models"
	"github.com/rostikts/fintech_test_project/test_utils"
	"math/rand"
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
		ID:              uint(rand.Intn(123456)),
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
		ID:              uint(rand.Intn(123456)),
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
	res, err := repo.GetRecords("")
	if err != nil {
		t.Fatalf("The list of transactions are not found due to the error: %v", err)
	}
	for _, v := range res {

		// exclude various fields
		v.DateInput = data.DateInput
		v.DatePost = data.DatePost

		if reflect.DeepEqual(v, data) {
			return
		}
	}

	t.Errorf("Created transaction is not found in the result list")
}

func TestTransactionRepositoryGetRecordsWithFilter(t *testing.T) {
	data := models.Transaction{
		RequestID:       2,
		TerminalID:      uint(rand.Intn(123456)),
		PartnerObjectID: 4,
		Payment: models.Payment{
			Type:      fmt.Sprintf("type %v", uint(rand.Intn(123456))),
			Number:    "num",
			Narrative: fmt.Sprintf("test partial search %v", uint(rand.Intn(123456))),
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
		Status:             fmt.Sprintf("test status %v", uint(rand.Intn(123456))),
	}
	data2 := models.Transaction{
		RequestID:       2,
		TerminalID:      3,
		PartnerObjectID: 44,
		Payment: models.Payment{
			Type:      "type",
			Number:    "num",
			Narrative: "narrative",
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
		DateInput:          time.Now(),
		DatePost:           time.Now(),
		Status:             "unique status for assert",
	}
	if err := repo.SaveTransaction(data); err != nil {
		t.Fatalf("The transaction is not saved due to the error: %v", err)
	}
	if err := repo.SaveTransaction(data2); err != nil {
		t.Fatalf("The transaction is not saved due to the error: %v", err)
	}
	tests := []struct {
		name     string
		filters  string
		expected models.Transaction
	}{
		{
			name:     "terminal_id filtering",
			filters:  fmt.Sprintf("WHERE terminal_id=%v", data.TerminalID),
			expected: data,
		},
		{
			name:     "terminal_id filtering",
			filters:  fmt.Sprintf("WHERE terminal_id=%v", data.TerminalID),
			expected: data,
		},
		{
			name:     "status filtering",
			filters:  fmt.Sprintf("WHERE status='%v'", data.Status),
			expected: data,
		},
		{
			name:     "payment_type filtering",
			filters:  fmt.Sprintf("WHERE payment.type='%v'", data.Payment.Type),
			expected: data,
		},
		{
			name:     "date range filtering",
			filters:  fmt.Sprintf("WHERE date_input>='%v' AND date_post<='%v'", data.DateInput.Format("2006-01-02"), data.DatePost.Format("2006-01-02")),
			expected: data,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res, err := repo.GetRecords(tc.filters)
			if err != nil {
				t.Fatalf("The list of transactions are not found due to the error: %v", err)
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

			if !reflect.DeepEqual(res[0], tc.expected) {
				t.Errorf("Returned incorrect element\n expected:%v\ngot: %v", tc.expected, res[0])
			}
		})
	}

}
