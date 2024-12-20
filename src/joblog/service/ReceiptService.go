package service

import (
	"context"
	"strconv"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
)

//
// func UploadReceipt(fileKey string, userId int, jobId int){
// 	client := data.InitS3()
//
// }

func GroupExpenseReceipts(Expenses []models.Expense) (map[models.Expense] models.Receipt) {
		ExpenseMap := dao.GetReceiptsByExpenseList(Expenses)
		return ExpenseMap
}

func DownloadReceipt(expId int) error {
	key, err := dao.GetReceiptKeyByExpenseId(expId)
	if err != nil {
		return err
	}
	client := data.InitS3()
	fileKey := strconv.Itoa(key)
	path := "receipts/" + fileKey
	data.DownloadS3(client, fileKey)
	presigner := data.InitS3PresignClient(client)
	data.GetObject(presigner, context.TODO(), "jobcontracts", path, 60)


	return nil

}
