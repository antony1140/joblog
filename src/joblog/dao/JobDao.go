package dao

import (
	"log"

	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
)

func CreateJob(job *models.Job){
	sql := "INSERT INTO job (title, description) values (?, ?)"	
	db := data.OpenDb()
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(job.Title, job.Description); err != nil {
		log.Fatal(err)
	}
}

func GetJobById(id int)(*models.Job, error){
	sql := "SELECT id, title, description, client_id FROM job WHERE id = ?"
	db := data.OpenDb()
	defer db.Close()
	var job models.Job
	stmt, err := db.Prepare(sql)
	if err != nil {
		return &job, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)

	if err := row.Scan(&job.Id, &job.Title, &job.Description, &job.ClientId); err != nil{
		return &job, err
	}
	return &job, err
}

func GetAllJobsByOrgId(id int)([]models.Job, error){
	sql := "SELECT id, title, description FROM job WHERE org_id = ?"
	var jobs []models.Job
	db := data.OpenDb()
	defer db.Close()
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
		if err:= rows.Scan(&job.Id, &job.Title, &job.Description); err != nil{
			return jobs, err
		}
		jobs = append(jobs, job)

	}
	return jobs, nil
}

func GetAllJobsByUserId(id int)([]models.Job, error){
	sql := "select job.id, job.title, job.Description from job join orgusers on job.org_id = orgusers.org_id where orgusers.user_id = ?"
	db := data.OpenDb()
	defer db.Close()
	var jobs []models.Job
	rows, err := db.Query(sql, id)
	if err != nil {
		log.Print(err)
		return jobs, err
	}

	for rows.Next(){
		var job models.Job
		if err := rows.Scan(&job.Id, &job.Title, &job.Description); err != nil {
			log.Print(err)
			return jobs, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil

}

func UpdateJob(job *models.Job)(int, error){
	sql := "UPDATE job SET title = ?, description = ?, contract = ?"
	db := data.OpenDb()
	defer db.Close()
	_, err := db.Exec(sql, job.Title, job.Description)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

// func UploadContract(org *models.Org, file io.Reader){
//
// 	client := data.InitS3()
// 	data.UploadS3(client, file)
//
// }
