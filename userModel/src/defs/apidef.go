package defs

import uuid "github.com/satori/go.uuid"

//requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//response
type SingledUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

// Data model
type VideoInfo struct {
	//视频id
	Id uuid.UUID
	//用户id
	AuthorId int
	//视频名字
	Name string
	//前端时间
	DisplayTime string
	//创建时间
	CreateTime string
}

// comments Model
type Comment struct {
	//评论id
	Id string
	//视频id
	VideoId string
	//用户id
	AuthorName string
	//评论内容
	Content string
	//评论时间
	ContentTime string
}

//session model
type SimpleSession struct {
	/**
	用户名称login_name
	*/
	UserName string
	/**
	超时时间
	*/
	TTL int64
}
