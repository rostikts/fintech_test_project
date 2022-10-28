package service

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/db/models"
	"github.com/rostikts/fintech_test_project/internal/transaction"
	"github.com/rostikts/fintech_test_project/pkg/datatypes"
	"io"
	"net/http"
	"os"
	"strings"
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
	fileName, err := downloadDocument(url)
	if err != nil {
		return 0, 0, err
	}

	// remove file to avoid the OOM exception
	defer func() {
		if err := os.Remove(fileName); err != nil {
			log.DefaultLogger.Error().Err(err).Msg("the tmp dump file is not deleted")
		}
	}()

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return 0, 0, err
	}

	parser := csv.NewReader(file)
	headers, err := parser.Read()
	if err != nil {
		return 0, 0, nil
	}

	wg := sync.WaitGroup{}
	for {
		tr, err := parseDocument(headers, parser)
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
	formattedFilter := prepareFilters(filters)
	res, err := s.repo.GetRecords(formattedFilter)
	if err != nil {
		return []models.Transaction{}, err
	}
	return res, nil
}

func downloadDocument(url string) (string, error) {
	fileName := fmt.Sprintf("%v.csv", uuid.New().String())
	output, err := os.Create(fileName)
	if err != nil {
		log.DefaultLogger.Error().Err(err).Msg("error occurred during file creation")
		return "", err
	}

	response, err := http.Get(url)
	if err != nil {
		log.DefaultLogger.Error().Err(err).Str("url", url).Msg("Error while downloading")
		return "", err
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func parseDocument(headers []string, reader *csv.Reader) (parsedTransaction, error) {
	record, err := reader.Read()
	if err != nil {
		return parsedTransaction{}, err
	}

	// small hack for unmarshalling single csv row into the struct
	tmpCsv := fmt.Sprintf("%v\n%v", strings.Join(headers, ","), strings.Join(record, ","))
	var tr []parsedTransaction
	if err := gocsv.UnmarshalString(tmpCsv, &tr); err != nil {
		return parsedTransaction{}, err
	}
	return tr[0], err
}

func prepareFilters(filters map[string]string) string {
	if len(filters) == 0 {
		return ""
	}
	result := "WHERE "
	counter := 0
	for k, v := range filters {
		switch k {
		case "terminal_id":
			result += fmt.Sprintf("%s=%s", k, v)
		case "transaction_id":
			result += fmt.Sprintf("t.id=%s", v)
		case "status":
			result += fmt.Sprintf("%s='%s'", k, v)
		case "payment_type":
			result += fmt.Sprintf("payment.type='%s'", v)
		case "from":
			result += fmt.Sprintf("date_input>='%v'", v)
		case "to":
			result += fmt.Sprintf("date_post<='%v'", v)
		case "payment_narrative":
			result += fmt.Sprintf("payment.narrative LIKE '%%%v%%'", v)

		}

		counter += 1
		if len(filters) != counter {
			result += " AND "
		}
	}
	return result
}
