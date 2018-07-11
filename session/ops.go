package session

type SimpleSession struct {
	Username string
	TTL      int64
}

var sessionMap *sync.Map


