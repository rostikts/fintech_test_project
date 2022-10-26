package models

import "time"

type Transaction struct {
	ID                 uint
	RequestID          uint
	TerminalID         uint
	PartnerObjectID    uint
	PaymentID          Payment
	ServiceID          Service
	PayeeID            Payee
	AmountTotal        int64
	AmountOriginal     int64
	CommissionPS       float64
	CommissionClient   int64
	CommissionProvider float64
	DateInput          time.Time
	DatePost           time.Time
	Status             string
}
