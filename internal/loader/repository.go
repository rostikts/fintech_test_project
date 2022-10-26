package loader

import "github.com/rostikts/fintech_test_project/models"

type Repository interface {
	SaveTransaction(data models.Transaction) error
	GetRecords() ([]models.Transaction, error)
}
