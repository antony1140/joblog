package dao

import (
	"github.com/antony1140/joblog/data"
)

func GetSessionUserBySid(sid string)(bool, int){
	sql := "select user_Id from session where sid = ?"
	db := data.OpenDb()
	defer db.Close()
	var userid int
	row := db.QueryRow(sql, sid)
	err := row.Scan(userid)
	if err != nil {
		return false, 0
	}
	
	return true, userid


}
