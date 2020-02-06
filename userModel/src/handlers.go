package main

import (
	"encoding/json"
	"github.com/bruce121381/juece-lol/src/dbops"
	"github.com/bruce121381/juece-lol/src/defs"
	"github.com/bruce121381/juece-lol/src/session"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

/**
创建用户
*/
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//post请求 读取请求体
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	/**
	序列化失败
	*/
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	//解析成功 但数据库返回失败
	if err := dbops.InsertUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
	//写入session
	sessionId := session.GenerateNewSession(ubody.Username)
	//返回消息体
	su := defs.SingledUp{Success: true, SessionId: sessionId}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
