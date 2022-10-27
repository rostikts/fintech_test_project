package transaction

import (
	"github.com/rostikts/fintech_test_project/db/models"
)

type Repository interface {
	SaveTransaction(data models.Transaction) error
	GetRecords() ([]models.Transaction, error)
}
