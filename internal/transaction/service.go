package transaction

import "github.com/rostikts/fintech_test_project/db/models"

type Service interface {
	ParseDocument(url string) (successCount, failedCount int64, err error)
	GetTransactions(filters map[string]string) ([]models.Transaction, error)
	GetTransactionsCSV(filters map[string]string) (string, error)
}
