package dao

import (
	"log"

	"github.com/antony1140/joblog/data"
	"github.com/antony1140/joblog/models"
)

func GetClientById(id int)(*models.Client, error){
	var client models.Client
	sql := "SELECT * from client where id = ?"

	db := data.OpenDb()
	row := db.QueryRow(sql, id)
	err := row.Scan(&client.Id, &client.Name, &client.ContactPref, &client.ContactSec)
	if err != nil {
		log.Print(err)
		return &client, err
	}
	log.Print("got a client: ", client.Id)

	return &client, nil
}
