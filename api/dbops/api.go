package dbops

import (
	"log"
	"database/sql"
	"github.com/yangruiyou85/video/api/defs"
	"github.com/yangruiyou85/video/utils"
	"time"
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

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {

	//create uuid

	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	//createtime-->db-->

	t := time.Now()
	ctime := t.Format("Jan 02 2006,15:04:05") //M D y，HH:MM:SS

	stmtIns, err := dbConn.Prepare(`insert into video_info(video_id,author_id,name,display_ctime)
                                value(?,?,?,?)`)

	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, nil

}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {


	stmtOut, err := dbConn.Prepare("select author_id,name,display_ctime from video_info where video_id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil

}

func DeleteVideoInfo(vid string) error {

	stmtDel, err := dbConn.Prepare("delete from video_info where video_id =?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil

}

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT a.vedio_id, a.author_id, a.name, a.display_ctime 
             FROM video_info a 		  
              JOIN users b ON a.author_id = b.users.id
		WHERE b.login_name = ? AND a.create_time > FROM_UNIXTIME(?) AND a.create_time <= FROM_UNIXTIME(?) 
		ORDER BY a.create_time DESC`)

	var res []*defs.VideoInfo
	if err != nil {
		return res, err

	}

	rows, err := stmtOut.Query(uname, from, to)
	if err != nil {
		log.Printf("%s", err)
		return res, err

	}
	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}

		vi := &defs.VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCtime: ctime}

		res = append(res, vi)
	}

	defer stmtOut.Close()
	return res, nil
}
