package types

import (
	"encoding/json"
	"testing"
)

func TestPasswordJSON(t *testing.T) {
	var u struct {
		Name     string   `json:"name"`
		Password Password `json:"password"`
	}

	bs := []byte(`{"name":"hello", "password":"asdfadsf"}`)
	err := json.Unmarshal(bs, &u)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", u.Password)
	if u.Password != "asdfadsf" {
		t.Errorf("u: %+v != asdfadsf", u)
		return
	}

	bs, err = json.Marshal(u)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("u bytes: %s", bs)
	err = json.Unmarshal(bs, &u)
	if err != nil {
		t.Error(err)
		return
	}
	if u.Password != "*******" {
		t.Errorf("u: %#v != *******", u)
		return
	}

}

func TestPassword(t *testing.T) {
	p1 := Password("asdfasdf")
	pg, err := p1.Generate()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%s", pg)

	err = pg.Compare([]byte("asdfasdf"))
	if err != nil {
		t.Error(err)
		return
	}
}
