package dao

import (
	"fmt"
	"log"

	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/security"
)

func GetUserByLogin(un string, pwd string) (*models.User, error){
	var user models.User
	pw := security.SecureCreds(pwd)
	sql := "select id, name, username, email from user where username = ? and pass = ?"
	db := data.OpenDb()
	defer db.Close()

	row := db.QueryRow(sql, un, pw)
	log.Println("un: " + un + ", pw: " + pw)
	
	if err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Email); err != nil{
		fmt.Println("err at userdao", err)	
		return &user, err
	}

	return &user, nil
}

func GetUserById(id int) (*models.User, error){
	var user models.User
	sql := "select name, username, email from user where id = ?"
	db := data.OpenDb()
	defer db.Close()

	row := db.QueryRow(sql, id)
	
	if err := row.Scan(&user.Name, &user.Username, &user.Email); err != nil{
		fmt.Println("err at userdao")	
		return &user, err
	}

	return &user, nil
}

func CreateUser(user *models.User)(int, error){
	sql := "insert into user (name, username, email, pass) values (?, ?, ?, ?)"

	db := data.OpenDb()
	defer db.Close()
	pw := user.Password
	pwd := security.SecureCreds(pw)
	fmt.Println(pwd)
	result, err := db.Exec(sql, &user.Name, &user.Username, &user.Email, pwd)
	if  err != nil {
		log.Println("couldnt create user", err)
		return 0, nil
	}
	r, rowErr := result.LastInsertId()
	log.Println("created user: ", int(r))
	return int(r), rowErr
}


