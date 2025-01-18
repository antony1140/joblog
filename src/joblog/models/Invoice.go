package models

type Invoice struct {
	Id int `json:"id"`
	Key string `json:"key"`
	Amount float32 `json:"amount"`
	Paid bool `json:"paid"`
	JobId int `json:"jobId"`
	RecipientName string `json:"recipientName"`
	RecipientContact string `json:"recipientContact"`
	RecipientEmail string `json:"recipientEmail"`
	RecipientAddress string `json:"recipientAddress"`
}

type InvoiceDTO struct {
	Amount float32 `json:"amount"`
	RecipientName string `json:"recipientName"`
	RecipientContact string `json:"recipientContact"`
	RecipientEmail string `json:"recipientEmail"`
	RecipientAddress string `json:"recipientAddress"`
	Paid bool `json:"paid"`
	JobId int `json:"job_id"`
}

type Recipient struct {
	Name string `json:"name"`
	Contact string `json:"contact"`
	Email string `json:"email"`
	Address string `json:"address"`
}

type InvoiceRequest struct {
	Recipient Recipient `json:"recipient"`
	//key: id, value: quantity
	Expenses map[string]string `json:"expenses"`
	JobId string `json:"jobId"`

}


func NewInvoice(amount float32, jobId int, recipientId int, paid bool) *Invoice {
	return &Invoice{
		Amount: amount,
		JobId: jobId,
		Paid: paid,
	}
}
