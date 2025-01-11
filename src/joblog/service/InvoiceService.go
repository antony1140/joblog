package service

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"github.com/antony1140/joblog/data"
)

func UploadInvoiceFromTempl(file string, invData interface{}) error {
	r, w := io.Pipe()
	tmpl, err := template.ParseGlob("views/invoice.html")
	if err != nil {
		log.Println(err)
		return err
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
	reader := bytes.NewReader(dat)

	client := data.InitS3() 
	data.UploadS3(client,reader, "invoice/" + file)

	return nil

}
