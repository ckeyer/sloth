package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Project 构建项目
type Project struct {
	ID            bson.ObjectId `json:"id" bson:"_id"`
	OwnerID       bson.ObjectId `json:"owner_id" bson:"owner_id"` // project owner
	Owner         string        `json:"owner" bson:"owner"`       // repos owner
	Name          string        `json:"name" bson:"name"`
	Desc          string        `json:"description" bson:"description"`
	ReposID       int64         `json:"repos_id" bson:"repos_id"`
	ReposName     string        `json:"repos_name" bson:"repos_name"`
	ReposFullname string        `json:"repos_fullname" bson:"repos_fullname"`
	Members       []Member      `json:"members" bson:"members"`
	Created       time.Time     `json:"created" bson:"created"`
}

// Member 项目用户，用于权限管理
type Member struct {
	UserID bson.ObjectId `json:"user_id" bson:"user_id"`
	Role   string        `json:"role" bson:"role"`
}

// CommitPatch one commit
type CommitPatch struct {
	SHA       string        `json:"sha" bson:"sha"`
	ProjectID bson.ObjectId `json:"project_id" bson:"project_id"`
	Author    string        `json:"author" bson:"author"`
	Message   string        `json:"message" bson:"message"`
}

// FilePatch file changes
type FilePatch struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	ProjectID bson.ObjectId `json:"project_id" bson:"project_id"`
	CommitSHA string        `json:"commit_sha" bson:"commit_sha"`
	Filename  string        `json:"filename" bson:"filename"`
	Mode      string        `json:"mode" bson:"mode"`
	Old       *LinePatch    `json:"old" bson:"old"`
	New       *LinePatch    `json:"new" bson:"new"`
}

type LinePatch struct {
	Offset int `json:"offset" bson:"offset"`
	Count  int `json:"count" bson:"count"`
}
