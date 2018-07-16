package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type NewComment struct {
	AuthorId int    `json:"author_id"`
	Content  string `json:"content"`
}

type NewVideo struct {
	AuthorId int    `json:"author_id"`
	Name     string `json:"name"`
}

type SignedUp struct {
	Sucess    bool   `json:"success"`
	SessionId string `json:"session_id"`
}

//data model

type VideoInfo struct {
	Id           string `json:"id"`
	AuthorId     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type VideosInfo struct {
	Videos []*VideoInfo `json:"videos"`
}

type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

type SimpleSession struct {
	Username string
	TTL      int64
}

type User struct {
	Id        int
	LoginName string
	Pwd       string
}

type UserInfo struct {
	Id int `json:"id"`
}
type UserSession struct {
	Username  string `json:"user_name"`
	SessionId string `json:"session_id"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

type SignedIn struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}
