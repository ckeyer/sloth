package types

import (
	"gopkg.in/mgo.v2/bson"
)

// Project 构建项目
type Project struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Type    string        `json:"type" bson:"type"`
	Members []Member      `json:"members" bson:"members"`
}

// 项目用户，用于权限管理
type Member struct {
	UserID bson.ObjectId `json:"user_id" bson:"user_id"`
	Role   string        `json:"role" bson:"role"`
}
