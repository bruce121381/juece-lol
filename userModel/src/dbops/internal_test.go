package dbops

import "testing"

func TestSelectAllSessions(t *testing.T) {
	sessions, err2 := SelectAllSessions()
	if err2 != nil {
		t.Errorf("%s", err2)
	}
	t.Log(sessions)
}
