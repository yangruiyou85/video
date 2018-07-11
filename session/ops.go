package session

import (
	"sync"
	"github.com/yangruiyou85/video/api/dbops"
	"github.com/yangruiyou85/video/api/defs"
	"github.com/yangruiyou85/video/utils"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap := &sync.Map{}
}

func LoadSessionFromDB() string {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true

	})
}

func GenerateNewSessionId(un string) string {

	id, _ := utils.NewUUID()
	ct := time.Now().UnixNano() / 1000000
	ttl := ct + 30*60*1000
	ss:=&defs.SimpleSession{Username:un,TTL:ttl}
	sessionMap.Store(id,ss)
	dbops.InsertSession(id,ttl,un)
	return id
}

func IsSessionExpired(sid string) (string, bool) {

}
