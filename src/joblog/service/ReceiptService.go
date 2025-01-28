package service

import (
	"context"
	"io"
	"log"
	"strconv"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

//
// func UploadReceipt(fileKey string, userId int, jobId int){
// 	client := data.InitS3()
//
// }

type request struct {
	req v4.PresignedHTTPRequest
}


func NewReceipt(receipt *models.Receipt, file io.Reader) (error) {
	client := data.InitS3()

	db := data.OpenDb()
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = dao.CreateReceipt(tx, receipt.FileKey, receipt.ExpenseId)

	if err != nil {
		tx.Rollback()
		return err
	}

	storageKey := "receipts/" + strconv.Itoa(receipt.ExpenseId) + "/" + receipt.FileKey
	err = data.UploadS3(client, file, storageKey)
	if err != nil {
		return err
	}

	err = tx.Commit() 
	return nil
}


func GroupExpenseReceipts(Expenses []*models.Expense) (map[*models.Expense] *models.Receipt) {
	ExpenseMap := dao.GetReceiptsByExpenseList(Expenses)
	for expense, receipt := range ExpenseMap {
		if receipt.Id != 0 {

			url, err := DownloadReceipt(expense.Id)
			if err != nil {
				log.Print(err)
				continue
			}
			receipt.S3Url = url.URL
		}
	}
	return ExpenseMap
}

func DownloadReceipt(expId int)(*v4.PresignedHTTPRequest, error)  {
	fileKey, err := dao.GetReceiptKeyByExpenseId(expId)
	if err != nil {
		log.Println(err)
	}
	client := data.InitS3()
	expIdStr:= strconv.Itoa(expId)
	path := "receipts/" + expIdStr + "/" + fileKey
	presigner := data.InitS3PresignClient(client)
	request, s3Err := data.GetObject(presigner, context.TODO(), "jobcontracts", path, 6)
	if s3Err != nil {
		log.Println("s3 failed to retreive presigned link", s3Err)
	}


	return request, err

}


func DownloadReceiptByFileKey(key string, expId int)(*v4.PresignedHTTPRequest, error)  {
	client := data.InitS3()
	expIdStr:= strconv.Itoa(expId)
	path := "receipts/" + expIdStr + "/" + key
	presigner := data.InitS3PresignClient(client)
	request, err := data.GetObject(presigner, context.TODO(), "jobcontracts", path, 6)
	if err != nil {
		log.Println("s3 failed to retreive presigned link", err)
	}


	return request, err

}

func DeleteReceipt(expId int)(error) {
	fileKey, err := dao.GetReceiptKeyByExpenseId(expId)
	if err != nil {
		log.Println(err)
	}
	client := data.InitS3()
	expIdStr:= strconv.Itoa(expId)
	path := "receipts/" + expIdStr + "/" + fileKey
	deleteErr := data.DeleteS3("jobcontracts", path, client, context.TODO())
	log.Print("trying to delete,", path)
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}
