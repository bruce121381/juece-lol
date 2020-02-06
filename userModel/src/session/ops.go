package session

import (
	"github.com/bruce121381/juece-lol/src/dbops"
	"github.com/bruce121381/juece-lol/src/defs"
	uuid "github.com/satori/go.uuid"
	"log"
	"sync"
	"time"
)

//存放session的缓存map
var sessionMap *sync.Map

//系统启动的时候将数据库中的session缓存到map中
func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

//加载session到缓存
func LoadSessionsFromDB() {
	sessions, err := dbops.SelectAllSessions()
	if err != nil {
		log.Printf("dbAllSessions is failed err is %s", err)
		return
	}

	sessions.Range(func(key, value interface{}) bool {
		session := value.(*defs.SimpleSession)
		sessionMap.Store(key, session)
		return true
	})
}

//产生一个新的session
func GenerateNewSession(userName string) string {
	id, _ := uuid.NewV4()
	ctime := nowInMilli()
	ttl := ctime + 5*60*1000

	simpleSession := &defs.SimpleSession{UserName: userName, TTL: ttl}
	sessionMap.Store(id, simpleSession)

	dbops.InsertSession(id, ttl, userName)

	return id.String()
}

//判断当前sessionId是否过期
func IsSessionExpired(sessionId string) (string, bool) {
	simpleSession, ok := sessionMap.Load(sessionId)
	if ok {
		ctime := nowInMilli()
		if simpleSession.(*defs.SimpleSession).TTL < ctime {
			//删除session
			deleteExpireSession(sessionId)
			return "", true
		}
		return simpleSession.(*defs.SimpleSession).UserName, false
	}
	return "", true

}

func deleteExpireSession(sessionId string) {
	sessionMap.Delete(sessionId)
	dbops.DeleteSession(sessionId)
}
