package admin

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id    bson.ObjectId `json:"_id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`

	Role     string `json:"role" bson:"role"` // admin member
	Password []byte `json:"-"`                // This is after encryption
}

type UserAuth struct {
	Id      bson.ObjectId `json:"_id"`
	UserId  bson.ObjectId `json:"user_id"`
	Token   []byte        `json:"token"`
	Expired time.Time     `json:"expired"`
}

type GithubAccount struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Token    string `bson:"token"`
}

type WebhookConfig struct {
	RepoName string `bson:""`
	Secret   string `bson:"secret"`
}

func (g *GithubAccount) GetToken() string {
	if g.Token != "" {
		return g.Token
	}
	return g.Password
}
