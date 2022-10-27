package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rostikts/fintech_test_project/internal/loader"
	"github.com/rostikts/fintech_test_project/models"
)

type transactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) loader.Repository {
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

	_, err = tx.Exec(`INSERT INTO transaction (request_id, terminal_id, partner_object_id, payment_id, service_id, payee_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status) 
							VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
		data.RequestID, data.TerminalID, data.PartnerObjectID, data.Payment.ID, data.Service.ID, data.Payee.ID, data.AmountTotal, data.AmountOriginal, data.CommissionPS, data.CommissionClient, data.CommissionProvider, data.DateInput, data.DatePost, data.Status)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (r transactionRepository) GetRecords() ([]models.Transaction, error) {
	var res []models.Transaction
	rows, err := r.db.Queryx(
		`SELECT t.id, t.request_id, t.terminal_id, t.partner_object_id, t.amount_original, t.amount_original, t.amount_total, t.commission_client, t.commission_client, t.commission_provider, t.commission_ps, t.date_input, t.date_post, t.status,
       				  p.id "payee.id", p.name "payee.name", p.bank_mfo "payee.bank_mfo", p.bank_account "payee.bank_account",
       				  s.id "service.id", s.name "service.name",
       				  p2.id "payment.id", p2.type "payment.type", p2.number "payment.number", p2.narrative "payment.narrative"
				FROM transaction t
    				INNER JOIN payee p on p.id = t.payee_id
    				INNER JOIN service s on s.id = t.service_id
    				INNER JOIN payment p2 on p2.id = t.payment_id;`)
	if err != nil {
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
