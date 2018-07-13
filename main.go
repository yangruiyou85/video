package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/yangruiyou85/video/utils"
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
	return router

}

func main() {

	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}

// handler->validation{}

//main--->middleware-defs(message,err)-->handlers--->dbops--->response--->
