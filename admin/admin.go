package admin

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id    bson.ObjectId `json:"_id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`

	Password []byte `json:"-"` // This is after encryption
}

type UserAuth struct {
	Id      bson.ObjectId `json:"_id"`
	UserId  bson.ObjectId `json:"user_id"`
	Token   []byte        `json:"token"`
	Expired time.Time     `json:"expired"`
}
