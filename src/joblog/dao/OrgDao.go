package dao

import (
	"fmt"
	"log"

	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
)

func CreateOrg(org *models.Org)(int, error){
	sql := "INSERT INTO Org (name, description) values (?, ?)"
	db := data.OpenDb()
	defer db.Close()
	result, err := db.Exec(sql, org.Name, org.Description)
	if err != nil {
		return 0, err
	}
	newId, resultErr := result.LastInsertId()
	if resultErr != nil {
		log.Print(resultErr)
	}
	return int(newId), nil

}

func AddOrgUser(userId int, orgId int)(error){
	sql := "INSERT INTO OrgUsers (user_id, org_id) values (?, ?)"
	db := data.OpenDb()
	defer db.Close()
	_, err := db.Exec(sql, userId, orgId)
	if err != nil {
		return err
	}
	return nil
}

func GetOrg(name string)(*models.Org, error){
	sql := "SELECT * FROM Job WHERE name = ?"
	db := data.OpenDb()
	defer db.Close()
	row := db.QueryRow(sql, name)
	var org models.Org
	if err := row.Scan(org.Id, org.Name); err != nil{
		return &org, err
	}
	return &org, nil
}
func GetOrgById(id int)(*models.Org, error){
	sql := "SELECT id, name, description FROM Org WHERE id = ?"
	db := data.OpenDb()
	defer db.Close()
	row := db.QueryRow(sql, id)
	var org models.Org
	if err := row.Scan(&org.Id, &org.Name, &org.Description); err != nil{
		return &org, err
	}
	return &org, nil
}

func GetOrgByJobId(id int)(*models.Org, error){
	sql := "select org.id, org.name, org.description from org join job on job.org_id = org.id where job.id = ?"	
	db := data.OpenDb()
	defer db.Close()
	row := db.QueryRow(sql, id)
	var org models.Org
	if err := row.Scan(&org.Id, &org.Name, &org.Description); err != nil{
		return &org, err
	}
	return &org, nil
}

func GetAllOrgsByUserId(id int) ([]models.Org, error) {
	orgs := []models.Org{}
	db := data.OpenDb()
	defer db.Close()
	sql := "select org.id, org.name from org join orgusers on org.id = orgusers.org_id where orgusers.user_id = ?;"

	rows, rowErr := db.Query(sql, id)
	if rowErr != nil {
		return orgs, rowErr
	}
	for rows.Next(){
		var org models.Org
		if err:= rows.Scan(&org.Id, &org.Name); err != nil{
			fmt.Println("there was an error in sql")
			return orgs, err
		}
		orgs = append(orgs, org)
		// fmt.Print("from sql org id: ", org.Id)
		fmt.Println(len(orgs))

	}
	return orgs, nil
}

func UpdateOrg(org *models.Org)(error){
	sql := "UPDATE Org SET name = ? WHERE id = ?"
	db := data.OpenDb()
	defer db.Close()
	_, err := db.Exec(sql, org.Name, org.Id)
	if err != nil {
		return err
	}
	return nil
}


