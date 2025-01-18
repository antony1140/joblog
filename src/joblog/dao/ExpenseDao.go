package dao

import (
	"errors"
	"log"

	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
)

func GetExpenseById(id int)(*models.Expense, error){
	sql := "select id, name, cost, job_id, description from expense where id = ?"
	var expense models.Expense

	db := data.OpenDb()
	defer db.Close()
	row := db.QueryRow(sql, id)
	err := row.Scan(&expense.Id, &expense.Name, &expense.Cost, &expense.JobId, &expense.Description)

	if err != nil {
		log.Print(err)
		return &expense, err
	}


	return &expense, nil

}

func GetAllExpensesByJobId(id int)([]*models.Expense, error){
	sql := "Select id, name, cost, description from expense where job_id = ?"
	var expenseList []*models.Expense

	db := data.OpenDb()
	defer db.Close()
	rows, err := db.Query(sql, id)
	if err != nil {
		log.Print(err)
		return expenseList, err
	}
	
	for rows.Next(){
		var expense models.Expense
		scanErr := rows.Scan(&expense.Id, &expense.Name, &expense.Cost, &expense.Description)
		if scanErr != nil {
			log.Print(scanErr)
			return expenseList, err
		}
		expenseList = append(expenseList, &expense)
	}

	return expenseList, nil

}


func CreateExpense(expense *models.Expense)(int, error){
	sql := "insert into expense (name, cost, job_id, description) values (?, ?, ?, ?)"
	db := data.OpenDb()
	defer db.Close()
	result, err := db.Exec(sql, expense.Name, expense.Cost, expense.JobId, expense.Description)
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

func UpdateExpenseById(expense *models.Expense)(int, error){
	sql := "update expense set name = ?, cost = ? where id = ?"
	db := data.OpenDb()
	defer db.Close()
	result, err := db.Exec(sql, expense.Name, expense.Cost, expense.Id)
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
		log.Print("rows affected in expenseDao ", rowsAffected)
		err := errors.New("something went wrong in updating the expsense")
		log.Print(err)
		return 0, err
	}

	return int(rowsAffected), nil
	

}

