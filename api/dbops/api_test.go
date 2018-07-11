package dbops

import (
	"testing"
	"strconv"
	"time"
	"fmt"
)

//    init(dblogin,truncate tables) ->run test->clear data(truncate table)

var tempvid string

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

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)

}

func testAddVideoInfo(t *testing.T) {

	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("error of AddVideoInfo:%v", err)
	}
	tempvid = vi.Id

}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("error of GetVideoInfo: %v", err)
	}

}

func testDeleteVideoInfo(t *testing.T) {

	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("error of DeleteVideoInfo:%v", err)
	}

}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("error of regetVideoInfo:%v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", TestListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "douyin"
	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("error of AddComments: %v", err)
	}
}

func TestListComments(t *testing.T) {

	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano() / 1000000000,10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("error of ListComments:%v", err)
	}
	for i, ele := range res {
		fmt.Printf("comment:%d,%v\n", i, ele)

	}

}
