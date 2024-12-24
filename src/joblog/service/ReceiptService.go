package service

import (
	"context"
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

func GroupExpenseReceipts(Expenses []models.Expense) (map[*models.Expense] *models.Receipt) {
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
	data.DownloadS3(client, fileKey)
	presigner := data.InitS3PresignClient(client)
	request, s3Err := data.GetObject(presigner, context.TODO(), "jobcontracts", path, 6)
	if s3Err != nil {
		log.Println("s3 failed to retreive presigned link", s3Err)
	}


	return request, err

}
