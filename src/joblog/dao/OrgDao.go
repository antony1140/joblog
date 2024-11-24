package dao

import (
	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
	"fmt"
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

func GetAllOrgsByUserId(id int) ([]models.Org, error) {
	orgs := []models.Org{}
	db := data.OpenDb()
	sql := "select name from org left outer join orgusers on org.id = orgusers.user_id where orgusers.user_id = ?;"

	rows, rowErr := db.Query(sql, id)
	if rowErr != nil {
		return orgs, rowErr
	}
	for rows.Next(){
		var org models.Org
		if err:= rows.Scan(&org.Name); err != nil{
			fmt.Println("there was an error in sql")
			return orgs, err
		}
		orgs = append(orgs, org)
		fmt.Println(len(orgs))

	}
		for org := range orgs {
			fmt.Println(org)
			fmt.Println("what even")
		}
	return orgs, nil
}

func UpdateOrg(org *models.Org)(error){
	sql := "UPDATE Org SET name = ? WHERE id = ?"
	db := data.OpenDb()
	_, err := db.Exec(sql, org.Name, org.Id)
	if err != nil {
		return err
	}
	return nil
}


