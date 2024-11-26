package security

import (
	"crypto/sha256"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"github.com/antony1140/joblog/data"
	"log"
	
)

func SecureCreds(pwd string)(string){
	pw := sha256.Sum256([]byte(pwd))	

	// user, err := dao.GetUserByLogin(un, string(pw[:]))

	// if err != nil{
	// 	return "", err	
	// }

	return string(pw[:])
}

func AuthenticateSession(){

}

func GetSession(c echo.Context)(bool, int){
	cookie, err := c.Cookie("sid")
	if err != nil {
			log.Print("got to cookie err")
		return false, 0
	}
	// user := dao.GetSessionUserBySid()
	sql := "select user_Id from sessions where sid = ?"
	db := data.OpenDb()
	var userid int
	row := db.QueryRow(sql, cookie.Value)
	scanErr := row.Scan(&userid)
	if scanErr != nil {
		log.Print(scanErr)
		
			log.Print("got to scan err")
		return false, 0
	}
		log.Print("session found for user", userid, "session: ", cookie.Value)
	return true, userid	
}

func CreateSession(insert bool, id int) *http.Cookie{
	sid := uuid.NewString()
	cookie := http.Cookie {
		Name:   "sid",
		Value: sid,
	// Quoted bool // indicates whether the Value was originally quoted
	//
	// Path       string    // optional
	// Domain     string    // optional
		Expires:   time.Now().Add(time.Second * 300),
	// RawExpires string    // for reading cookies only
	//
	// // MaxAge=0 means no 'Max-Age' attribute specified.
	// // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// // MaxAge>0 means Max-Age attribute present and given in seconds
	// MaxAge      int
	// Secure      bool
		HttpOnly:    true,
	// SameSite    SameSite
	// Partitioned bool
	// Raw         string
	// Unparsed    []string // Raw text of unparsed attribute-value pairs
		
	}
	if insert {
	sql := "insert into sessions (sid, user_id) values (?, ?)"
	db := data.OpenDb()
	_, err := db.Exec(sql, sid, id)	
	if err != nil{
		log.Print(err)
	}

	}
	return &cookie
}

