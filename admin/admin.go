package admin

import (
	"time"

	"github.com/ckeyer/sloth/types"
	"gopkg.in/mgo.v2/bson"
)

type User types.User

type UserAuth struct {
	Id      bson.ObjectId `json:"_id"`
	UserId  bson.ObjectId `json:"user_id"`
	Token   []byte        `json:"token"`
	Expired time.Time     `json:"expired"`
}
