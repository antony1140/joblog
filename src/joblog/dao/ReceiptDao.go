package dao

import (
	"database/sql"
	"log"

	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
)

func CreateReceipt(tx *sql.Tx, fileName string, expenseId int) (int, error){
	sql := "INSERT INTO receipt(expense_id, fileKey) values (?, ?)"
	// defer db.Close()

	result, err := tx.Exec(sql, expenseId, fileName )
	if err != nil {
		log.Print("err at receipt dao: create,",  err)
		return 0, err
	}

	id, err := result.LastInsertId()	
	tx.Commit()


	return int(id), nil
}

func GetReceiptKeyByExpenseId(expId int) (string, error){
	sql := "SELECT fileKey from receipt where expense_id = ?"
	db := data.OpenDb()
	defer db.Close()
	row := db.QueryRow(sql, expId)
	var receipt models.Receipt
	err := row.Scan(&receipt.FileKey)
	if err != nil {
		return "", err
	}
	return receipt.FileKey, nil
}

func GetReceiptById(id int)(*models.Receipt, error) {
	var receipt models.Receipt
	sql := "Select id, expense_id, fileKey from receipt where id = ?"
	db := data.OpenDb()
	defer db.Close()
	
	result := db.QueryRow(sql, id)
	err := result.Scan(&receipt.Id, &receipt.ExpenseId, &receipt.FileKey)

	return &receipt, err
}

func GetReceiptsByExpenseList(expenses []*models.Expense)(map[*models.Expense] *models.Receipt){
	sql := "Select id, expense_id, fileKey from receipt where expense_id = ?"
	db := data.OpenDb()
	defer db.Close()
	var receipts []models.Receipt
	ExpenseMap := make(map[*models.Expense] *models.Receipt)
	for _, expense := range expenses {
		var receipt models.Receipt
		row := db.QueryRow(sql, expense.Id)
		err := row.Scan(&receipt.Id, &receipt.ExpenseId, &receipt.FileKey)
		if err != nil {
			receipt.Id = 0
			ExpenseMap[expense] = &receipt
			continue
		}
		receipts = append(receipts, receipt)	
		ExpenseMap[expense] = &receipt
	}

	return ExpenseMap
}

func GetReceiptsByExpense(expense *models.Expense)([]*models.Receipt, error) {
	var list []*models.Receipt
	sql := "select id, expense_id, fileKey from receipt where expense_id = ?"
	db := data.OpenDb()
	defer db.Close()
	rows, err := db.Query(sql, expense.Id)

	for rows.Next() {
		var receipt models.Receipt
		row := db.QueryRow(sql, expense.Id)
		err = row.Scan(&receipt.Id, &receipt.ExpenseId, &receipt.FileKey)
		if err != nil {
			continue
		}
		list = append(list, &receipt)	
	}


	return list, err
}

func DeleteReceiptById(id int)(error) {
	sql := "delete from receipt where id = ?"
	db := data.OpenDb()
	defer db.Close()
	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	
	return nil
}

func DeleteReceiptByExpenseId(id int)(error) {
	sql := "delete from receipt where expense_id = ?"
	db := data.OpenDb()
	defer db.Close()
	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	
	return nil
}

