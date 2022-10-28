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

func TestPrepareFilters(t *testing.T) {
	tests := []struct {
		name     string
		filters  map[string]string
		expected string
	}{
		{
			name:     "Numeral and non numeral filter",
			filters:  map[string]string{"transaction_id": "32", "status": "accepted"},
			expected: "WHERE t.id=32 AND status='accepted'",
		},
		{
			name:     "Non numeral filter",
			filters:  map[string]string{"payment_type": "test"},
			expected: "WHERE payment.type='test'",
		},
		{
			name:     "Date filter",
			filters:  map[string]string{"from": "2021-02-10"},
			expected: "WHERE date_input>='2021-02-10'",
		},
		{
			name:     "Narrative filter",
			filters:  map[string]string{"payment_narrative": "teest її"},
			expected: "WHERE payment.narrative LIKE '%teest її%'",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := prepareFilters(tc.filters)
			if tc.expected != result {
				t.Errorf("The filters are prepared incorrectly\nExpected: %s\nActual: %s", tc.expected, result)
			}
		})
	}
}
