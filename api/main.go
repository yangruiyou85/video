package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/yangruiyou85/video/api/session"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {

	m := middleWareHandler{}
	m.r = r

	return m

}

func (m middleWareHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {

	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {

	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	//router.GET("/user/:username", GetUserInfo)
	router.POST("/user/:username/videos", AddNewVideo)

	router.GET("/user/:username/videos", ListAllVideos)

	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)

	router.POST("/videos/:vid-id/comments", PostComment)

	router.GET("/videos/:vid-id/comments", ShowComments)
	return router

}

func Prepare() {

	session.LoadSessionFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}

// handler->validation{}

//main--->middleware-defs(message,err)-->handlers--->dbops--->response--->
