package dbops

import (
	"database/sql"
	"github.com/bruce121381/juece-lol/src/defs"
	uuid "github.com/satori/go.uuid"
	"log"
	"strconv"
	"sync"
)

/**
session与db交互的class
*/

//写数据库session
func InsertSession(sessionId uuid.UUID, ttl int64, userName string) error {
	ttlInt := strconv.FormatInt(ttl, 10)
	stmtIns, err2 := dbConn.Prepare("INSERT into user_session (session_id,TTL,login_name) VALUES (?,?,?)")
	if err2 != nil {
		return err2
	}
	_, err2 = stmtIns.Exec(sessionId, ttlInt, userName)
	if err2 != nil {
		return err2
	}
	defer stmtIns.Close()
	return nil
}

//获取session
func SelctSession(sessionId string) (*defs.SimpleSession, error) {
	simpleSession := &defs.SimpleSession{}
	stmtOut, err2 := dbConn.Prepare("SELECT TTL,login_name from user_session where session_id = ?")
	if err2 != nil {
		return nil, err2
	}
	var ttl string
	var userName string
	err2 = stmtOut.QueryRow(sessionId).Scan(&ttl, &userName)
	if err2 != nil && err2 != sql.ErrNoRows {
		return nil, err2
	}

	if res, err2 := strconv.ParseInt(ttl, 10, 64); err2 != nil {
		simpleSession.TTL = res
		simpleSession.UserName = userName
	} else {
		return nil, err2
	}
	defer stmtOut.Close()

	return simpleSession, nil
}

//删除session
func DeleteSession(sessionId string) error {
	stmtDel, err2 := dbConn.Prepare("DELETE FROM user_session where session_id  = ?")
	if err2 != nil {
		log.Printf("dbConnSession is failed, err is %s", err2)
		return err2
	}
	_, err2 = stmtDel.Exec(sessionId)
	if err2 != nil {
		log.Printf("dbDelSession is failed, err is %s", err2)
		return err2
	}

	defer stmtDel.Close()
	return nil

}

//获取sessionList
func SelectAllSessions() (*sync.Map, error) {
	sessionMap := &sync.Map{}
	stmtOut, err2 := dbConn.Prepare("select * from user_session")
	if err2 != nil {
		log.Printf("dbconn error ,err is %s", err2)
		return nil, err2
	}
	rows, err2 := stmtOut.Query()
	if err2 != nil {
		log.Printf("dbQuery error, err is %s", err2)
		return nil, err2
	}

	for rows.Next() {
		var sessionId string
		var ttlStr string
		var userName string

		if err := rows.Scan(&sessionId, &ttlStr, &userName); err != nil {
			log.Printf("select sessionList error, err is %s", err)
			break
		}

		if ttl, err1 := strconv.ParseInt(ttlStr, 10, 64); err1 != nil {
			simpleSession := &defs.SimpleSession{UserName: userName, TTL: ttl}
			sessionMap.Store(sessionId, simpleSession)
			log.Printf("session id is %s ,ttl is %d", sessionId, ttl)
		}

	}

	defer stmtOut.Close()
	return sessionMap, nil

}
