package models

import (
	// "github.com/antony1140/joblog/data"
	"strconv"

)

type Job struct {
	Id int
	Title string
	Description string
	ClientId int
	Contract bool
	OrgId int
}

func NewJob(title string, desc string) *Job{
	return &Job {
		Title: title,
		Description: desc,
	}
}

func PrintJob (job *Job)(string){
	return "job{ " + strconv.Itoa(job.Id) + ",\n " + job.Title + ",\n " + job.Description + "\n}"
}



