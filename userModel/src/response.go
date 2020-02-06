package main

import (
	"encoding/json"
	"github.com/bruce121381/juece-lol/src/defs"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	//将错误码写入响应头中
	w.WriteHeader(errResp.HttpSC)
	//json反序列化
	resStr, _ := json.Marshal(&errResp.Error)
	//响应头响应
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {

	//将正确码写入响应头
	w.WriteHeader(sc)
	//响应头响应
	io.WriteString(w, resp)

}
