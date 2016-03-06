package types

import (
	"github.com/ckeyer/go-ci/lib"
)

func init() {
	db := lib.InitDB()
	db.DB().Ping()
	db.CreateTable(&User{})
}
