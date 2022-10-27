package models

import "time"

type Transaction struct {
	ID                 uint      `db:"id" csv:"TransactionId"`
	RequestID          uint      `db:"request_id" csv:"RequestId"`
	TerminalID         uint      `db:"terminal_id" csv:"TerminalId"`
	PartnerObjectID    uint      `db:"partner_object_id" csv:"PartnerObjectId"`
	AmountTotal        int64     `db:"amount_total" csv:"AmountTotal"`
	AmountOriginal     int64     `db:"amount_original" csv:"AmountOriginal"`
	CommissionPS       float64   `db:"commission_ps" csv:"CommissionPS"`
	CommissionClient   int64     `db:"commission_client" csv:"CommissionClient"`
	CommissionProvider float64   `db:"commission_provider" csv:"CommissionProvider"`
	DateInput          time.Time `db:"date_input" csv:"-"`
	DatePost           time.Time `db:"date_post" csv:"-"`
	Status             string    `db:"status" csv:"Status"`
	Payment            Payment
	Service            Service
	Payee              Payee
}
