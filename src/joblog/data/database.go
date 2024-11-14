package data

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
	"os"
)

func InitDb(){
	os.Create("./data/joblog.db")
	db, err := sql.Open("sqlite3", "./data/joblog.db")
	if err != nil {
		log.Fatal(err)
	}


	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS Job (id INTEGER PRIMARY KEY AUTOINCREMENT, title varchar(255), description varchar(255), contract INTEGER, Foreign Key (contract) References Contracts (id));")

	statement.Exec()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}


func OpenDb() *sql.DB{
	db, err := sql.Open("sqlite3", "./data/joblog.db")
	if err != nil {
		log.Fatal(err)
	}
	return db

}
