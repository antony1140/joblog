package dao

import (
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/security"
	"fmt"
)

func GetUserByLogin(un string, pwd string) (*models.User, error){
	var user models.User
	pw := security.SecureCreds(pwd)
	sql := "select id, name, username from user where username = ? and pass = ?"
	db := data.OpenDb()

	row := db.QueryRow(sql, un, pw)
	
	if err := row.Scan(&user.Id, &user.Name, &user.Username); err != nil{
		fmt.Println("err at userdao")	
		return &user, err
	}

	return &user, nil
}

func CreateUser(user *models.User)(int, error){
	sql := "insert into user (name, username, pass) values (?, ?, ?)"

	db := data.OpenDb()
	pw := user.Password
	pwd := security.SecureCreds(pw)
	fmt.Println(pwd)
	result, err := db.Exec(sql, &user.Name, &user.Username, pwd)
	if  err != nil {
		return 0, nil
	}
	r, rowErr := result.LastInsertId()
	return int(r), rowErr
}


