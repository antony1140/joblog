package dao

import (
	"log"

	"github.com/antony1140/joblog/data"
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
