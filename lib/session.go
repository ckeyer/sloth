package lib

import (
	"github.com/martini-contrib/sessions"
)

func GetCookieStore() sessions.CookieStore {
	return sessions.NewCookieStore([]byte("cZig4BdxWHdse"))
}

// func GetSession() sessions.Session {
// 	return sessions.Sessions("kyt-api", GetCookieStore())
// }
