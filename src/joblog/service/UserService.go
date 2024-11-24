package service

import (
	"github.com/antony1140/joblog/models"
	"github.com/antony1140/joblog/dao"
)

func LoginUser(un string, pwd string) (*models.User, error){
	user, err := dao.GetUserByLogin(un, pwd)	
	if err != nil{
		return user, err
	}

	return user, nil
}

func CreateUser(user *models.User) (int, error){
	newId, err := dao.CreateUser(user)
	if err != nil {
		return newId, err
	}

	return newId, err
}

