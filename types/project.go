package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Project 构建项目
type Project struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	OwnerID   bson.ObjectId `json:"owner_id" bson:"owner_id"`
	Name      string        `json:"name" bson:"name"`
	Desc      string        `json:"description" bson:"description"`
	ReposID   int64         `json:"repos_id" bson:"repos_id"`
	ReposName string        `json:"repos_name" bson:"repos_name"`
	Members   []Member      `json:"members" bson:"members"`
	Created   time.Time     `json:"created" bson:"created"`
}

// 项目用户，用于权限管理
type Member struct {
	UserID bson.ObjectId `json:"user_id" bson:"user_id"`
	Role   string        `json:"role" bson:"role"`
}
