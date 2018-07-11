package dbops

import "testing"

//    init(dblogin,truncate tables) ->run test->clear data(truncate table)

func clearTables() {
	dbConn.Exec("truncate table users;")
	dbConn.Exec("truncate table sessions;")
	dbConn.Exec("truncate table video_info;")
	dbConn.Exec("truncate table comments;")
}

func TestMain(m *testing.M) {

	clearTables()
	m.Run()
	clearTables()

}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", TestDeleteUserUser)
	t.Run("Reget", testRegetUser)

}

func testAddUser(t *testing.T) {

	err := AddUserCredential("kkgo", "root")
	if err != nil {
		t.Errorf("error of adduser:%v", err)
	}

}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("kkgo")
	if pwd != "root" || err != nil {
		t.Errorf("error of getUser:%v", err)
	}

}

func TestDeleteUserUser(t *testing.T) {
	err := DeleteUser("kkgo", "root")
	if err != nil {
		t.Errorf("error of DeleteUser:%v", err)
	}

}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("kkgo")
	if err != nil {
		t.Errorf("error of RegetUser:%v", err)
	}

	if pwd != "" {
		t.Errorf("erroe of pwd:%v", err)
	}

}
