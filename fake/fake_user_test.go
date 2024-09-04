package fake

import "testing"

func TestUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		u, e := GenerateUsrInfo(3)
		if e != nil {
			t.Error(e)
		}
		for j := range u {
			t.Logf("%d generated %+v", i+1, u[j])
		}
	}
}
