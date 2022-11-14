package service

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/config"
	"github.com/rostikts/fintech_test_project/db/models"
	"github.com/rostikts/fintech_test_project/internal/transaction"
	"github.com/rostikts/fintech_test_project/pkg/datatypes"
	"io"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
)

type parsedTransaction struct {
	models.Transaction
	DateInput datatypes.DateTime `db:"date_input" csv:"DateInput"`
	DatePost  datatypes.DateTime `db:"date_post" csv:"DatePost"`
}

func (tr parsedTransaction) ToModel() models.Transaction {
	res := tr.Transaction
	res.DatePost = tr.DatePost.Time
	res.DateInput = tr.DateInput.Time
	return res
}

type transactionService struct {
	repo transaction.Repository
}

func NewTransactionService(repository transaction.Repository) transaction.Service {
	return transactionService{repo: repository}
}

func (s transactionService) ParseDocument(url string) (successCount, failedCount int64, err error) {
	body, err := s.downloadDocument(url)

	reader := csv.NewReader(body)

	if err != nil {
		return 0, 0, err
	}
	unmarshaler, err := gocsv.NewUnmarshaller(reader, parsedTransaction{})
	if err != nil {
		return 0, 0, err
	}

	wg := sync.WaitGroup{}
	for {
		tr, err := s.parseDocument(unmarshaler)
		if err == io.EOF {
			break
		}
		wg.Add(1)
		go func() {
			if err := s.SaveTransaction(tr.ToModel()); err != nil {
				log.DefaultLogger.Error().Err(err).Msg("The document is not saved to db")
				atomic.AddInt64(&failedCount, 1)
			}
			atomic.AddInt64(&successCount, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	return
}

func (s transactionService) SaveTransaction(transaction models.Transaction) error {
	err := s.repo.SaveTransaction(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (s transactionService) GetTransactions(filters map[string]string) ([]models.Transaction, error) {
	formattedFilters, arguments := prepareFilters(filters)
	res, err := s.repo.GetRecords(formattedFilters, arguments...)
	if err != nil {
		return []models.Transaction{}, err
	}
	return res, nil
}

func (s transactionService) GetTransactionsCSV(filters map[string]string) (string, error) {
	transactions, err := s.GetTransactions(filters)
	if err != nil {
		return "", err
	}
	fileName, err := s.storeTransactionsInFile(transactions)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func (s transactionService) storeTransactionsInFile(transactions []models.Transaction) (string, error) {
	fileName := fmt.Sprintf("%v/%v.csv", config.Config.FilePath, uuid.New().String())
	output, err := os.Create(fileName)
	if err != nil {
		log.DefaultLogger.Error().Err(err).Msg("error occurred during file creation")
		return "", err
	}
	if err := gocsv.MarshalFile(transactions, output); err != nil {
		return "", err
	}
	return fileName, nil
}

func (s transactionService) downloadDocument(url string) (io.ReadCloser, error) {

	response, err := http.Get(url)
	if err != nil {
		log.DefaultLogger.Error().Err(err).Str("url", url).Msg("Error while downloading")
		return nil, err
	}

	if response.StatusCode > 299 {
		return nil, fmt.Errorf("invalid status code received during download")
	}
	return response.Body, nil
}

func (s transactionService) parseDocument(unmarshaler *gocsv.Unmarshaller) (parsedTransaction, error) {
	res, err := unmarshaler.Read()
	if err != nil {
		return parsedTransaction{}, err
	}
	if tr, ok := res.(parsedTransaction); ok {
		return tr, nil
	}

	return parsedTransaction{}, fmt.Errorf("the row was not parsed properly")
}

func prepareFilters(filters map[string]string) (string, []interface{}) {
	if len(filters) == 0 {
		return "", []interface{}{}
	}
	arguments := make([]interface{}, len(filters))
	result := "WHERE "
	counter := 0
	for k, v := range filters {
		position := counter + 1
		switch k {
		case "terminal_id":
			result += fmt.Sprintf("t.terminal_id=$%v", position)
		case "status":
			result += fmt.Sprintf("%s=$%v", k, position)
		case "transaction_id":
			result += fmt.Sprintf("t.id=$%v", position)
		case "payment_type":
			result += fmt.Sprintf("payment.type=$%v", position)
		case "from":
			result += fmt.Sprintf("date_input>=$%v", position)
		case "to":
			result += fmt.Sprintf("date_post<=$%v", position)
		case "payment_narrative":
			result += fmt.Sprintf("payment.narrative LIKE $%v", position)
			v = "%" + v + "%"

		}
		arguments[counter] = v
		counter += 1
		if len(filters) != counter {
			result += " AND "
		}
	}
	return result, arguments
}
