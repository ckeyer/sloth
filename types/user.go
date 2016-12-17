package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	RoleAdmin  = "admin"
	RoleMember = "member"
	RoleGuest  = "guest"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Phone    string        `json:"phone" bson:"phone"`
	Password Password      `json:"password" bson:"password"`
	Role     string        `json:"role" bson:"role"`

	Created   time.Time `json:"created" bson:"created"`
	Updated   time.Time `json:"updated" bson:"updated"`
	LastLogin time.Time `json:"last_login" bson:"last_login"`

	GithubAccount *GithubAccount `json:"github,omitempty" bson:"github,omitempty"`
	WechatAccount *WechatAccount `json:"wechat,omitempty" bson:"wechat,omitempty"`
}

type WechatAccount struct {
	Name string `json:"name" bson:"name"`
}

type GithubAccount struct {
	Name      string   `json:"name" bson:"name"`
	Email     string   `json:"email" bson:"email"`
	Password  Password `json:"password" bson:"password"`
	DeployKey Password `json:"deploy_key" bson:"deploy_key"`
	Token     Password `json:"token" bson:"token"`
}

type UserAuth struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	UserId  bson.ObjectId `json:"user_id" bson:"user_id"`
	Token   string        `json:"token" bson:"token"`
	Created time.Time     `json:"created" bson:"created"`
	Lasted  time.Time     `json:"lasted" bson:"lasted"`
	Expired time.Time     `json:"expired" bson:"expired"`
}

// func (g *GithubAccount) GetToken() string {
// 	if len(g.Token) == 0 {
// 		return string(g.Token)
// 	}
// 	return string(g.Password)
// }
