package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/security"
	"github.com/antony1140/joblog/service"
	"github.com/labstack/echo/v4"
)

func NewInvoice(c echo.Context) (error) {
	user, _ := security.GetSession(c)
	if !user {
		return c.Redirect(302, "/")
	}


	var invoiceRequest models.InvoiceRequest
	err := json.NewDecoder(c.Request().Body).Decode(&invoiceRequest)
	if err != nil {
		log.Println("err at decode", err)
		return c.NoContent(500)

	}
	fmt.Println("decoded: \n", invoiceRequest)

	invoice, expenseList, err := service.AggregateInvoice(&invoiceRequest.Recipient, invoiceRequest.Expenses, invoiceRequest.JobId)
	if err != nil {
		return c.NoContent(500)
	}

	invoice.Key = "invoice" + strconv.Itoa(invoice.Id) + ".html"
	_, err = dao.UpdateInvoice(invoice)
	if err != nil {
		log.Println("dao err", err)
	}
	fileTypes := strings.Split(invoice.Key, ".")
	htmlKey := fileTypes[0] + ".html"
	pdfKey := fileTypes[0] + ".pdf"


	log.Println("expense map in question: ", expenseList)

	invData := struct {
		ExpenseList map[*models.Expense] int
		Invoice *models.Invoice

	} {
		ExpenseList: expenseList,
		Invoice: invoice,
	}

	_, err = service.UploadInvoiceFromTempl(invoice.Key, invData)
	if err != nil {
		log.Println(err)
		return c.HTML(500, "<h1> something went wrong </h1>")
	}

	file,_ := os.Open("./" + pdfKey)
	client := data.InitS3()
	err = data.UploadS3(client, file, "invoice/" + pdfKey)
	if err != nil {
		log.Println(err)
		return c.HTML(500, "something broke at upload")
	}


	invoice.Key = "invoice" + pdfKey
	dao.UpdateInvoice(invoice)
	err = data.DeleteS3("invoice/" + htmlKey, client, context.TODO())
	if err != nil {
		log.Println(err)
		return c.HTML(500, "something broke at delete")
	}
	os.Remove("./" + pdfKey)



	jobId := strconv.Itoa(invoice.JobId)
	return c.Redirect(302, "/job/" + jobId)
}
