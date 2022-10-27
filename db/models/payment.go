package models

type Payment struct {
	ID        uint   `csv:"-"`
	Type      string `csv:"PaymentType"`
	Number    string `csv:"PaymentNumber"`
	Narrative string `csv:"PaymentNarrative"`
}
