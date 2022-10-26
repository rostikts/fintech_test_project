package repository

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/config"
	"github.com/rostikts/fintech_test_project/models"
	"github.com/rostikts/fintech_test_project/test_utils"
	"os"
	"reflect"
	"testing"
	"time"
)

var repo transactionRepository

func TestMain(m *testing.M) {
	err := godotenv.Load("./../../../.env")
	if err != nil {
		log.DefaultLogger.Fatal().Msg("Error loading .env file" + err.Error())
	}
	db := test_utils.NewTestDB(config.Config.DBConfig)
	defer db.Close()

	dns := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Config.DBConfig.User,
		config.Config.DBConfig.Password,
		"localhost",
		4011,
		config.Config.DBConfig.Name,
	)
	migration, err := migrate.New("file://./../../../db/migrations", dns)
	if err != nil {
		log.DefaultLogger.Fatal().Err(err).Msg("error occurred during the creation of the migrations")
	}

	err = migration.Down()
	if err != nil {
		log.DefaultLogger.Error().Err(err).Msg("error occurred")
	}

	if err := migration.Up(); err != nil {
		if err.Error() != "no change" {
			log.DefaultLogger.Fatal().Err(err).Msg("error occurred")
		}
	}
	repo = transactionRepository{db: db}
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
