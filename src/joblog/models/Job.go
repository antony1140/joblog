package models

import (
	// "github.com/antony1140/joblog/data"

)

type Job struct {
	Id int
	Title string
	Description string
	Contract int
}

func NewJob(title string, desc string) *Job{
	var job Job
	job.Title = title
	job.Description = desc
	return &job
}

func GetId(job *Job) int {
	return job.Id
}
func GetTitle(job *Job) string {
	return job.Title
}
func GetDesc(job *Job) string {
	return job.Description
}
