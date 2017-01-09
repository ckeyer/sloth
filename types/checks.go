package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type CheckResp struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	ProjectID  bson.ObjectId `json:"project_id" bson:"project_id"`
	SHA1       string        `json:"sha1" bson:"sha1"`
	CommitID   string        `json:"commit_id" bson:"commit_id"`
	Filename   string        `json:"filename" bson:"filename"`
	CheckType  string        `json:"check_type" bson:"check_type"`
	CheckName  string        `json:"check_name" bson:"check_name"`
	LineNumber int           `json:"line" bson:"line"`
	Error      string        `json:"error" bson:"error"`
	Created    time.Time     `json:"created" bson:"created"`
}
