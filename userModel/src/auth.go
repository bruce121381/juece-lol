package main

import (
	"net/http"
	"userModel/src/session"
)

var HEADER_FILED_SESSION = "X-Session-Id"
var HEADER_FILED_UNAME = "X-User-Name"

/**
验证session的合法性
*/
func validateSession(r *http.Request) bool {
	sessionId := r.Header.Get(HEADER_FILED_SESSION)
	if len(sessionId) == 0 {
		return false
	}
	userName, ok := session.IsSessionExpired(sessionId)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FILED_UNAME, userName)
	return true
}

/**
验证用户登录的合法性
*/
func vaildateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FILED_UNAME)
	if len(uname) == 0 {
		//sendErrorResponse(w)
		return false
	}
	return true
}
