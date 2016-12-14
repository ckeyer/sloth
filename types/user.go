package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Phone    string        `json:"phone" bson:"phone"`
	Password Password      `json:"password" bson:"password"`

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

// func (g *GithubAccount) GetToken() string {
// 	if len(g.Token) == 0 {
// 		return string(g.Token)
// 	}
// 	return string(g.Password)
// }
