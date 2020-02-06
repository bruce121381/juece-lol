package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type MiddleWareHandler struct {
	r *httprouter.Router
}

func (m MiddleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	validateSession(r)
	//过滤请求
	m.r.ServeHTTP(w, r)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := MiddleWareHandler{}
	m.r = r
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	//创建用户
	router.PUT("/user", CreateUser)
	//用户登录
	router.POST("/user/:user_name", Login)
	//删除用户

	//添加video

	//获取单个video信息

	//删除video

	//新增评论

	//获取某个视频的所有评论

	return router
}

func main() {
	r := RegisterHandlers()
	handler := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", handler)

}
