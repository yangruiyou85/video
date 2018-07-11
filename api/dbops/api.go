package dbops

import "log"

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

	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string ) (string, error) {

	stmtOut, err := dbConn.Prepare("select pwd from users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)

	stmtOut.Close()
	return pwd, nil

}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name=? and  pwd=?")
	if err != nil {
		log.Printf("Delete User error:%s", err)
		return err
	}

	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return err
}
