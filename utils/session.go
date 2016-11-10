package utils

import (
	"github.com/martini-contrib/sessions"
)

func GetCookieStore() sessions.CookieStore {
	return sessions.NewCookieStore([]byte("cZig4BdxWHdse"))
}
