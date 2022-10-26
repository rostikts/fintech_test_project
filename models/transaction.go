package models

import "time"

type Transaction struct {
	ID                 uint      `db:"id"`
	RequestID          uint      `db:"request_id"`
	TerminalID         uint      `db:"terminal_id"`
	PartnerObjectID    uint      `db:"partner_object_id"`
	AmountTotal        int64     `db:"amount_total"`
	AmountOriginal     int64     `db:"amount_original"`
	CommissionPS       float64   `db:"commission_ps"`
	CommissionClient   int64     `db:"commission_client"`
	CommissionProvider float64   `db:"commission_provider"`
	DateInput          time.Time `db:"date_input"`
	DatePost           time.Time `db:"date_post"`
	Status             string    `db:"status"`
	Payment            Payment
	Service            Service
	Payee              Payee
}
