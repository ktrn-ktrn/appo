package authentication

import (
	"database/sql"
	"sync"
)

var mx sync.Mutex
var sessions map[string]*Session

func Get(token string) *Session {
	mx.Lock()
	session, ok := sessions[token]
	if !ok {
		return nil
	}
	mx.Unlock()
	return session
}

func Set(token string, session *Session) {
	mx.Lock()
	sessions[token] = session
	mx.Unlock()
}

type Session struct {
	UserId    int64
	DBConnect *sql.DB
}

func Login() {

}

func Logout() {

}
