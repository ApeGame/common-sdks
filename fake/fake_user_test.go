package fake

import "testing"

func TestUser(t *testing.T) {
	u, e := GenerateUsrInfo()
	if e != nil {
		t.Error(e)
	}
	t.Log(u)
}
