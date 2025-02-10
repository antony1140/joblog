package data

import (
	"database/sql"
	"log"
	"os"
	"strings"
	_ "github.com/mattn/go-sqlite3"
)


func InitDb(){
	var src string
	if os.Getenv("MODE") == "prod" {
		src = "./data/prod_joblog.db"
	} else {
		src = "./data/joblog.db"
	}
	os.Open("./data/prod_joblog.db")
	db, err := sql.Open("sqlite3", src)
	defer db.Close()
	if err != nil {
		log.Println("err 1")
		log.Fatal(err)
	}

		schema := strings.Split(os.Getenv("DB_SCHEMA"), ";")
		for _, column := range schema {
			_, err := db.Exec(column)
			if err != nil {
				log.Println("error executing line: " + column + ", error:", err)
			}
		}
	}









func OpenDb() *sql.DB{
	var source string
	if os.Getenv("MODE") == "prod" {
		source = "./data/prod_joblog.db"	
	} else {
		source = "./data/joblog.db"
	}

	db, err := sql.Open("sqlite3", source)
	if err != nil {
		log.Fatal(err)
	}
	return db

}
