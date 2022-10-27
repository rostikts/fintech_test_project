package models

type Payee struct {
	ID          uint   `db:"id" csv:"PayeeId"`
	Name        string `db:"name" csv:"PayeeName"`
	BankMfo     uint   `db:"bank_mfo" csv:"PayeeBankMfo"`
	BankAccount string `db:"bank_account" csv:"PayeeBankAccount"`
}
