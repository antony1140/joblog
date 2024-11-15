package dao

import (
	"log"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/data"
	"io"
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

func GetJob(job *models.Job)(*models.Job, error){
	sql := "SELECT * FROM job WHERE id = ?"
	db := data.OpenDb()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return job, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(job.Id)
	if err := row.Scan(&job.Id, &job.Title, &job.Description, &job.Contract); err != nil{
		return job, err
	}
	return job, err
}

func GetAllJobsByOrg(id int)([]models.Job, error){
	sql := "SELECT * FROM job WHERE org = ?"
	var jobs []models.Job
	db := data.OpenDb()
	stmt, err := db.Prepare(sql)
	if err != nil {
		return jobs, err
	}
	defer stmt.Close()
	rows, rowErr := db.Query(sql, id)
	if rowErr != nil {
		return jobs, rowErr
	}
	for rows.Next(){
		var job models.Job
		if err:= rows.Scan(&job.Id, &job.Title, &job.Description, &job.Contract); err != nil{
			return jobs, err
		}
		jobs = append(jobs, job)

	}
	return jobs, nil
}

func UpdateJob(job *models.Job)(int, error){
	sql := "UPDATE job SET title = ?, description = ?, contract = ?"
	db := data.OpenDb()
	_, err := db.Exec(sql, job.Title, job.Description, job.Contract)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func UploadContract(org *models.Org, file io.Reader){
	
	client := data.InitS3()
	data.UploadS3(client, file)

}
