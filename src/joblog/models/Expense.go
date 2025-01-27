package models

type Expense struct {
	Id int
	JobId int
	Name string
	Cost string
	Quant int
	InvoiceQty int
	Description string
	Receipt bool
	AmountPaid float32
	DocList []*Receipt
}


