package dbops

import (
	"log"
	"database/sql"
)

//func OpenConn() *sql.DB {
//
//	dbConn, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/video?charset=utf8")
//	if err != nil {
//		panic(err.Error())
//
//	}
//	return dbConn
//
//
//}

func AddUserCredential(loginName string, pwd string) error {

	stmtIns, err := dbConn.Prepare("insert into users(login_name,pwd)values(?,?)")
	if err != nil {
		return err

	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {

	stmtOut, err := dbConn.Prepare("select pwd from users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()
	return pwd, nil

}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name=? and  pwd=?")
	if err != nil {
		log.Printf("Delete User error:%s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	stmtDel.Close()
	return err
}
