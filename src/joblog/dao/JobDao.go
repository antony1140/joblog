package dao

import (
	"log"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/data"

)

func CreateJob(job *models.Job){
	sql := "INSERT INTO job (title, description) values (?, ?)"	
	db := data.OpenDb()
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(job.Title, job.Description); err != nil {
		log.Fatal(err)
	}
}
