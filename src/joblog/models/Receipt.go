package models

type Receipt struct {
	Id int
	ExpenseId int
	FileKey string
	S3Url string
}
