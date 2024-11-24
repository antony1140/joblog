package security

import (
	"crypto/sha256"
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

