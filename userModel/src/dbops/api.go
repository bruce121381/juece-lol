package dbops

import (
	"database/sql"
	"github.com/bruce121381/juece-lol/src/defs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

//添加用户
func InsertUserCredential(loginName string, password string) error {

	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name,password,create_time,last_login_time,modified_time) VALUES (?,?,?,?,?)")

	if err != nil {
		return err
	}
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmtIns.Exec(loginName, password, timeNow, timeNow, timeNow)
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	return nil
}

//获取用户信息
func SelectUserCredential(loginName string) (string, error) {

	stmtout, err := dbConn.Prepare("select password from users where login_name = ? ")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var password string
	err = stmtout.QueryRow(loginName).Scan(&password)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtout.Close()
	return password, nil
}

//用户注销
func DeleteUser(loginName string, password string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name = ? and password = ?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, password)
	if err != nil {
		return nil
	}
	defer stmtDel.Close()
	return nil
}

//------------------------------------------------------------------------------------------------------------------------

//添加video
func InsertVideo(authorId int, name string) (*defs.VideoInfo, error) {

	//create uuid
	videoId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	dtime := time.Now().Format("2006-01-02 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO
		video_info (id,author_id,name,display_time,create_time) values (?,?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	cTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmtIns.Exec(videoId, authorId, name, dtime, cTime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: videoId, AuthorId: authorId, Name: name, DisplayTime: dtime, CreateTime: cTime}
	defer stmtIns.Close()
	return res, nil

}

//获取video信息
func SelectVideo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("select author_id,name,display_time from video_info where id = ?")

	var author_id int
	var dct string
	var name string
	err = stmtOut.QueryRow(vid).Scan(&author_id, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	uuidVid, _ := uuid.FromString(vid)
	res := &defs.VideoInfo{Id: uuidVid, AuthorId: author_id, DisplayTime: dct, Name: name}
	return res, nil
}

//删除video信息
func DeleteVideo(vid string) error {

	stmtDel, err := dbConn.Prepare("delete  FROM  video_info where id = ?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	return nil
}

//------------------------------------------------------------------------------------------------------------------------

//新增评论
func InsertComments(vid string, aid int, content string) error {
	//评论id
	cId, err2 := uuid.NewV4()
	if err2 != nil {
		return err2
	}
	stmtIns, err2 := dbConn.Prepare("INSERT INTO comments (id,video_id,author_id,content,content_time) VALUES (?,?,?,?,?)")
	if err2 != nil {
		return nil
	}
	_, err2 = stmtIns.Exec(cId, vid, aid, content, time.Now().Format("2006-01-02 15:04:05"))
	if err2 != nil && err2 != sql.ErrNoRows {
		return err2
	}
	if err2 == sql.ErrNoRows {
		return err2
	}
	//关闭链接
	defer stmtIns.Close()
	return nil
}

//查看评论
func SelectCommentList(vid string, from, to int) ([]*defs.Comment, error) {

	stmtOut, err2 := dbConn.Prepare(`
SELECT
	c.id,a.login_name,c.content
FROM
	comments AS c
	INNER JOIN users AS a ON a.id = c.author_id
WHERE
c.video_id = ? and c.content_time > FROM_UNIXTIME(?) and c.content_time <= FROM_UNIXTIME(?) 
`)

	if err2 != nil {
		return nil, err2
	}

	var res []*defs.Comment

	rows, err2 := stmtOut.Query(vid, from, to)
	if err2 != nil {
		return nil, err2
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		comments := &defs.Comment{Id: id, VideoId: vid, AuthorName: name, Content: content}
		res = append(res, comments)

	}
	defer stmtOut.Close()
	return res, nil
}
