package models

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Username string `json:"username"`
	Password string 
	Email string `json:"email"`
}

