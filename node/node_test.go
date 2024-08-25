package node

import "testing"

func TestGetChainIdByName(t *testing.T) {
	id, err := GetChainIdByName("bsc")
	if err != nil {
		t.Error(err)
	}
	t.Log(id)
}

func TestGetChainNameById(t *testing.T) {
	name, err := GetChainNameById("1")
	if err != nil {
		t.Error(err)
	}
	t.Log(name)
}
