package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/phuslu/log"
	"github.com/rostikts/fintech_test_project/db/models"
	"github.com/rostikts/fintech_test_project/internal/transaction"
)

type transactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) transaction.Repository {
	return transactionRepository{db: db}
}

func (r transactionRepository) SaveTransaction(data models.Transaction) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO service (id, name) VALUES ($1, $2)`, data.Service.ID, data.Service.Name)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`INSERT INTO payee (id, name, bank_mfo, bank_account) VALUES ($1, $2, $3, $4)`, data.Payee.ID, data.Payee.Name, data.Payee.BankMfo, data.Payee.BankAccount)
	if err != nil {
		return err
	}

	rows, err := tx.Query(`INSERT INTO payment (type, number, narrative) VALUES ($1, $2, $3) RETURNING id`, data.Payment.Type, data.Payment.Number, data.Payment.Narrative)
	defer rows.Close()
	if err != nil {
		return err
	}
	for rows.Next() {
		if err := rows.Scan(&data.Payment.ID); err != nil {
			return err
		}
	}

	_, err = tx.Exec(`INSERT INTO transaction 
    							(request_id, terminal_id, partner_object_id, payment_id, service_id, payee_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status) 
								VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
		data.RequestID,
		data.TerminalID,
		data.PartnerObjectID,
		data.Payment.ID,
		data.Service.ID,
		data.Payee.ID,
		data.AmountTotal,
		data.AmountOriginal,
		data.CommissionPS,
		data.CommissionClient,
		data.CommissionProvider,
		data.DateInput,
		data.DatePost,
		data.Status)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r transactionRepository) GetRecords(filters string) ([]models.Transaction, error) {
	var res []models.Transaction
	qry := fmt.Sprintf(`SELECT t.id, t.request_id, t.terminal_id, t.partner_object_id, t.amount_original, t.amount_original, t.amount_total, t.commission_client, t.commission_client, t.commission_provider, t.commission_ps, t.date_input, t.date_post, t.status,
       				  payee.id "payee.id", payee.name "payee.name", payee.bank_mfo "payee.bank_mfo", payee.bank_account "payee.bank_account",
       				  s.id "service.id", s.name "service.name",
       				  payment.id "payment.id", payment.type "payment.type", payment.number "payment.number", payment.narrative "payment.narrative"
				FROM transaction t
    				INNER JOIN payee on payee.id = t.payee_id
    				INNER JOIN service s on s.id = t.service_id
    				INNER JOIN payment  on payment.id = t.payment_id 
    			%s`, filters)

	rows, err := r.db.Queryx(qry)
	if err != nil {
		log.DefaultLogger.Error().Err(err).Msg(qry)
		return []models.Transaction{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var tr models.Transaction
		if err := rows.StructScan(&tr); err != nil {
			return []models.Transaction{}, err
		}
		res = append(res, tr)
	}
	return res, nil
}
