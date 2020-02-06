package dbops

import (
	uuid "github.com/satori/go.uuid"
	"testing"
)

/**
func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate user_session")
}

*/

/**
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

*/

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("delete", testDeleteUser)
	t.Run("ReGet", testReGetUser)
}

func testAddUser(t *testing.T) {

	err := InsertUserCredential("mxz-test", "123")
	if err != nil {
		t.Errorf("Error of AddUser exceptiong is %v", err)
	}

}

func testGetUser(t *testing.T) {
	pwd, err := SelectUserCredential("mxz-test")
	if err != nil {
		t.Error("Error of GetUser")
	}
	if pwd != "123" {
		t.Errorf("Error of password")
	}

}

func testDeleteUser(t *testing.T) {
	err2 := DeleteUser("mxz-test", "123")
	if err2 != nil {
		t.Errorf("Error of DeleteUser exceptiong is %v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err2 := SelectUserCredential("mxz-test")
	if pwd != "123" || err2 != nil {
		t.Error("Error of ReGetUser")
	}
}

//------------------------------------------------------------------------------------------------------------------------

//添加视频
func TestAddVideo(t *testing.T) {
	video, err2 := InsertVideo(3, "老人与海")
	if err2 != nil {
		t.Errorf("add User failed, err is %s", err)
	}
	t.Log(video)
}

//删除视频
func TestDeleteVideo(t *testing.T) {
	err2 := DeleteVideo("ca3bdf05-54aa-4881-a9c0-fee26538aed0")
	if err2 != nil {
		t.Errorf("delete Video failed, err is %s", err2)
	}
}

//查找视频
func TestGetVideo(t *testing.T) {
	video, err2 := SelectVideo("ca3bdf05-54aa-4881-a9c0-fee26538aed0")
	if err2 != nil {
		t.Errorf("get Video is failed,vid %s, err is %s", video.Id, err2)
	}
	t.Log(video)

}

//------------------------------------------------------------------------------------------------------------------------

//找出当前视频的所有评论
func TestGetComments(t *testing.T) {
	list, err2 := SelectCommentList("ca3bdf05-54aa-4881-a9c0-fee26538aed0", 1580486400, 1580529600)
	if err2 != nil {
		t.Errorf("Select Comments List is failed, err is %s", err2)
	}
	t.Log(list)
}

//新增评论
func TestInsertComments(t *testing.T) {
	v4, err2 := uuid.NewV4()
	if err2 != nil {
		t.Error("uuid failed")
	}
	err2 = InsertComments(v4.String(), 3, "好好看啊")
	if err2 != nil {
		t.Errorf("error of %s", err2)
	}
}
