package service

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
)

func AggregateInvoice(recipient *models.Recipient, expenses map[string]string, jobId string) (*models.Invoice, map[*models.Expense]int, error) {
	expList := make(map[*models.Expense]int)
	var newInvoice models.Invoice
	for id, qty := range expenses {
		expQty, err := strconv.Atoi(qty)
		if err != nil {
			log.Println("error aggregating invoice data", err)
			continue
		}
		expId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("error aggregating invoice data", err)
			continue
		}
		exp, err := dao.GetExpenseById(expId)	
		if err != nil {
			log.Println("error aggregating invoice data", err)
			continue
		}
		expList[exp] = expQty

		value, err := strconv.ParseFloat(exp.Cost, 32)
		if err != nil {
			log.Println("error aggregating invoice data", err)
			continue
		}
		expCost := float32(value) * float32(expQty)
		newInvoice.Amount += expCost
	}
	
	job, err := strconv.Atoi(jobId)	
	if err != nil {
		log.Println(err)
	}
	newInvoice.JobId = job
	newInvoice.RecipientName = recipient.Name
	newInvoice.RecipientAddress = recipient.Address
	newInvoice.RecipientContact = recipient.Contact
	newInvoice.RecipientEmail = recipient.Email
	newInvoice.Paid = false

	fmt.Println("aggregated invoice data: ", newInvoice)

	newId, err := dao.CreateInvoice(&newInvoice)
	if err != nil {
		log.Println("err creating an invoice", err)
	}
	newInvoice.Id = newId
	


	return &newInvoice, expList, err	
}

//using wkhtmltopdf
func UploadInvoiceFromTemplWK(file string, invData interface{}) (string, error) {
	r, w := io.Pipe()
	tmpl, err := template.ParseGlob("views/invoice.html")
	if err != nil {
		log.Println(err)
		return "", err
	}
	go func() {
		tmpErr := tmpl.ExecuteTemplate(w, "invoice", invData)
		if tmpErr != nil {
			log.Println("err in the go routine,", tmpErr)
		}

		w.Close()

	}()
	

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r)
	dat := buffer.Bytes()
	html := buffer.String()
	reader := bytes.NewReader(dat)

	client := data.InitS3() 
	data.UploadS3(client,reader, "invoice/" + file)

	return html, nil

}


func UploadInvoiceFromTempl(file string, invData interface{}) (string, error) {
	fileTypes := strings.Split(file, ".")
	htmlKey := fileTypes[0] + ".html"
	pdfKey := fileTypes[0] + ".pdf"
	r, w := io.Pipe()
	tmpl, err := template.ParseGlob("views/invoice.html")
	if err != nil {
		log.Println(err)
		return "", err
	}
	go func() {
		tmpErr := tmpl.ExecuteTemplate(w, "newInvoice", invData)
		if tmpErr != nil {
			log.Println("err in the go routine,", tmpErr)
		}

		w.Close()

	}()
	

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(r)
	dat := buffer.Bytes()
	html := buffer.String()
	reader := bytes.NewReader(dat)

	client := data.InitS3() 
	data.UploadS3(client,reader, "invoice/" + htmlKey)


	err = data.DownloadS3WithKey(client, file, "invoice/" + htmlKey)
	if err != nil {
		log.Println("err at s3 download", err)
		return "", err
	}

	cmd := exec.Command("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "--headless", "--disable-gpu", "--no-pdf-header-footer", "--print-to-pdf=./"+ pdfKey,"./assets/" + htmlKey)

	err = cmd.Run()
	if err != nil {
		log.Println("err at cmd run", err)
		return "", err
	}



	return html, nil

}
