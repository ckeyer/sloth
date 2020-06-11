package api

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/ckeyer/sloth/pkgs/admin"
	"github.com/ckeyer/sloth/pkgs/gh"
	"github.com/ckeyer/sloth/global"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ListProjects(ctx *gin.Context) {
	db := ctx.MustGet(CtxKeyMgoDB).(*mgo.Database)
	skip, limit := getListQuery(ctx)

	ret := []*gh.Project{}
	err := db.C(global.ColProject).Find(bson.M{}).Skip(skip).Limit(limit).All(&ret)
	if err != nil {
		log.Errorf("ListProjects, %s", err.Error())
		GinError(ctx, 500, err)
		return
	}

	ctx.JSON(200, ret)
}

func NewProject(ctx *gin.Context) {
	var p gh.Project
	err := json.NewDecoder(ctx.Request.Body).Decode(&p)
	if err != nil {
		GinError(ctx, 400, err)
		return
	}

	u := ctx.MustGet(CtxKeyUser).(*admin.User)
	p.ID = bson.NewObjectId()
	p.OwnerID = u.ID
	p.Created = time.Now()

	db := ctx.MustGet(CtxKeyMgoDB).(*mgo.Database)
	err = db.C(global.ColProject).Insert(&p)
	if err != nil {
		log.Errorf("NewProject, %s", err.Error())
		GinError(ctx, 500, err)
		return
	}

	ctx.JSON(201, p)
}

func GetProject(ctx *gin.Context) {
	id := ctx.Param("id")
	if !bson.IsObjectIdHex(id) {
		GinError(ctx, 400, "ivalid project id.")
		return
	}

	db := ctx.MustGet(CtxKeyMgoDB).(*mgo.Database)
	ret := &gh.Project{}
	err := db.C(global.ColProject).FindId(bson.ObjectIdHex(id)).One(ret)
	if err != nil {
		GinError(ctx, 500, err)
		return
	}

	ctx.JSON(200, ret)
}

func UpdateProject(ctx *gin.Context) {
	id := ctx.Param("id")
	if !bson.IsObjectIdHex(id) {
		GinError(ctx, 400, "ivalid project id.")
		return
	}

	var p gh.Project
	err := json.NewDecoder(ctx.Request.Body).Decode(&p)
	if err != nil {
		GinError(ctx, 400, err)
		return
	}

	db := ctx.MustGet(CtxKeyMgoDB).(*mgo.Database)
	update := bson.M{
		"$set": bson.M{
			"name":        p.Name,
			"description": p.Desc,
		},
	}
	err = db.C(global.ColProject).UpdateId(bson.ObjectIdHex(id), update)
	if err != nil {
		log.Errorf("UpdateProject. %s", err.Error())
		GinError(ctx, 500, err)
		return
	}

	err = db.C(global.ColProject).FindId(bson.ObjectIdHex(id)).One(&p)
	if err != nil {
		log.Errorf("UpdateProject. find latest one, %s", err.Error())
		GinError(ctx, 500, err)
		return
	}

	ctx.JSON(200, p)
}

func RemoveProject(ctx *gin.Context) {
	id := ctx.Param("id")
	if !bson.IsObjectIdHex(id) {
		GinError(ctx, 400, "ivalid project id.")
		return
	}

	db := ctx.MustGet(CtxKeyMgoDB).(*mgo.Database)
	err := db.C(global.ColProject).RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		log.Errorf("RemoveProject, %s", err.Error())
		GinError(ctx, 500, err)
		return
	}

	GinMessage(ctx, 200, "ok")
}
