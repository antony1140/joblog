package service

import (
	"database/sql"
	"log"

	"github.com/antony1140/joblog/dao"
	"github.com/antony1140/joblog/models"
)

type JobService struct {
	Db *sql.DB
}

func GetJobById(id int)(*models.Job, error){
	job, err := dao.GetJobById(id)	
	if err != nil {
		return nil, err
	}
	return job, nil
}

func NewJob(job *models.Job) (error) {
	err := dao.CreateJob(job)
	if err != nil {
		log.Println("error creating job at dao", err)
		return err
	}

	return nil
}
