package service

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/internal/loader"
	"github.com/rostikts/fintech_test_project/models"
	"github.com/rostikts/fintech_test_project/pkg/datatypes"
	"io/ioutil"
	"net/http"
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

type loaderService struct {
	repo loader.Repository
}

func NewLoaderService(repository loader.Repository) loader.Service {
	return loaderService{repo: repository}
}

func (s loaderService) ParseDocument(url string) (successCount, failedCount int64, err error) {
	document, err := downloadDocument(url)
	if err != nil {
		return 0, 0, err
	}
	parsedTrs, err := parseDocument(document)
	if err != nil {
		return 0, 0, err
	}
	for _, v := range parsedTrs {
		if err := s.SaveTransaction(v.ToModel()); err != nil {
			log.DefaultLogger.Error().Err(err).Msg("The document is not saved to db")
			failedCount += 1
		}
		successCount += 1
	}
	return
}

func (s loaderService) SaveTransaction(transaction models.Transaction) error {
	err := s.repo.SaveTransaction(transaction)
	if err != nil {
		return err
	}
	return nil
}

func downloadDocument(url string) ([]byte, error) {
	fmt.Println("Downloading file...")

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return []byte{}, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func parseDocument(data []byte) ([]parsedTransaction, error) {
	var result []parsedTransaction
	if err := gocsv.UnmarshalBytes(data, &result); err != nil {
		return []parsedTransaction{}, err
	}
	return result, nil
}
