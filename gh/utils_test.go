package gh

import (
	"gopkg.in/check.v1"
)

func (g *GhSuite) TestMustString(c *check.C) {
	str := "abc"
	num := 123
	var nilPtr *int
	var nilStr *string
	for k, v := range map[interface{}]string{
		"":     "",
		nil:    "",
		nilPtr: "",
		nilStr: "",
		"a":    "a",
		str:    "abc",
		&str:   "abc",
		num:    "123",
		&num:   "123",
	} {
		if ret := mustString(k); v != ret {
			c.Errorf("%#v should be: <%s>, but <%s>", k, v, ret)
		}
	}
}
