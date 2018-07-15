package session

import (
	"sync"
	"github.com/yangruiyou85/video/api/dbops"
	"github.com/yangruiyou85/video/api/defs"
	"github.com/yangruiyou85/video/api/utils"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000

}

func LoadSessionFromDB() {
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
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)
	return id
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)

}

func IsSessionExpired(sid string) (string, bool) {

	ss, ok := sessionMap.Load(sid)
	ct := nowInMilli()
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false

	} else {
		ss, err := dbops.RetrieveSession(sid)
		if err != nil || ss == nil {
			return "", true

		}
		if ss.TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		sessionMap.Store(sid, ss)
		return ss.Username, false
	}

	return "", true

}
