package models

type Payee struct {
	ID          uint   `db:"id"`
	Name        string `db:"name"`
	BankMfo     uint   `db:"bank_mfo"`
	BankAccount string `db:"bank_account"`
}
