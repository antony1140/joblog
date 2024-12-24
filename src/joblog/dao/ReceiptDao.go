package dao

import (
	"fmt"
	"log"

	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
)

func CreateReceipt(fileName string, expenseId int) (int, error){
	sql := "INSERT INTO receipt(expense_id, fileKey) values (?, ?)"
	db := data.OpenDb()

	result, err := db.Exec(sql, expenseId, fileName )
	if err != nil {
		log.Print("err at receipt dao: create,",  err)
		return 0, err
	}

	id, err := result.LastInsertId()	


	return int(id), nil
}

func GetReceiptKeyByExpenseId(expId int) (string, error){
	sql := "SELECT fileKey from receipt where expense_id = ?"
	db := data.OpenDb()
	row := db.QueryRow(sql, expId)
	var receipt models.Receipt
	err := row.Scan(&receipt.FileKey)
	if err != nil {
		return "", err
	}
	return receipt.FileKey, nil
}

func GetReceiptsByExpenseList(expenses []models.Expense)(map[*models.Expense] *models.Receipt){
	sql := "Select id, expense_id, fileKey from receipt where expense_id = ?"
	db := data.OpenDb()
	var receipts []models.Receipt
	ExpenseMap := make(map[*models.Expense] *models.Receipt)
	for _, expense := range expenses {
		fmt.Print("debug expense id", expense.Id, expense.Name)
		var receipt models.Receipt
		row := db.QueryRow(sql, expense.Id)
		err := row.Scan(&receipt.Id, &receipt.ExpenseId, &receipt.FileKey)
		if err != nil {
			log.Print("error at getReceipt ByExpenseList", err)
			receipt.Id = 0
			ExpenseMap[&expense] = &receipt
			continue
		}
		receipts = append(receipts, receipt)	
		ExpenseMap[&expense] = &receipt
	}

	return ExpenseMap
}
