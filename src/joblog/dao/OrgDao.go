package dao

import (
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
)

func CreateOrg(org *models.Org)(error){
	sql := "INSERT INTO Job (name) values (?)"
	db := data.OpenDb()
	_, err := db.Exec(sql, org.Name)
	if err != nil {
		return err
	}
	return nil

}

func GetOrg(name string)(*models.Org, error){
	sql := "SELECT * FROM Job WHERE name = ?"
	db := data.OpenDb()
	row := db.QueryRow(sql, name)
	var org models.Org
	if err := row.Scan(org.Id, org.Name); err != nil{
		return &org, err
	}
	return &org, nil
}

func UpdateOrg(org *models.Org)(error){
	sql := "UPDATE Org SET name = ?"
	db := data.OpenDb()
	_, err := db.Exec(sql, org.Name)
	if err != nil {
		return err
	}
	return nil
}


