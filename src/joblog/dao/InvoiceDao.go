package dao

import (
	"errors"
	"log"

	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
)


func GetInvoiceById(id int)(*models.Invoice, error){
	sql := "select id, key, job_id from invoice where id = ?"
	var invoice models.Invoice

	db := data.OpenDb()
	defer db.Close()
	row := db.QueryRow(sql, id)
	err := row.Scan(&invoice.Id, &invoice.Key, &invoice.JobId)

	if err != nil {
		log.Print(err)
		return &invoice, err
	}


	return &invoice, nil

}

func GetAllInvoicesByJobId(id int)([]*models.Invoice, error){
	sql := "Select id, key, job_id from invoice where job_id = ?"
	var invoiceList []*models.Invoice

	db := data.OpenDb()
	defer db.Close()
	rows, err := db.Query(sql, id)
	if err != nil {
		log.Print(err)
		return invoiceList, err
	}
	
	for rows.Next(){
		var invoice models.Invoice
		scanErr := rows.Scan(&invoice.Id, &invoice.Key, &invoice.JobId)
		if scanErr != nil {
			log.Print(scanErr)
			return invoiceList, err
		}
		invoiceList = append(invoiceList, &invoice)
	}

	return invoiceList, nil

}


func CreateInvoice(invoice *models.Invoice)(int, error){
	sql := "insert into invoice (key, job_id, recipient_name, recipient_email, recipient_address, paid, amount) values (?, ?, ?, ?, ?, ?, ?)"
	db := data.OpenDb()
	defer db.Close()
	result, err := db.Exec(sql, invoice.Key, invoice.JobId, invoice.RecipientName, invoice.RecipientEmail, invoice.RecipientAddress, invoice.Paid, invoice.Amount)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	newId, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
		return 0, err
	}

	return int(newId), nil
}


func UpdateInvoice(invoice *models.Invoice)(int, error){
	sql := "update invoice set key = ?, job_id = ? where id = ?"
	db := data.OpenDb()
	defer db.Close()
	result, err := db.Exec(sql, &invoice.Key, &invoice.JobId, &invoice.Id)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Print(err)
		return 0, err
	}

	if rowsAffected != 1 {
		err := errors.New("something went wrong in updating the invoice")
		log.Print(err)
		return 0, err
	}

	return int(rowsAffected), nil
	

}

