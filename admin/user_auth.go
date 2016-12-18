package admin

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/commons/util"
	"github.com/ckeyer/sloth/global"
	"github.com/ckeyer/sloth/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserAuth types.UserAuth

func NewUserAuth(uid bson.ObjectId, expired time.Time) *UserAuth {
	return &UserAuth{
		Id:      bson.NewObjectId(),
		UserId:  uid,
		Created: time.Now(),
		Lasted:  time.Now(),
		Token:   util.RandomString(15),
		Expired: expired,
	}
}

func AuthSignature(db *mgo.Database, apiKey, timestamp, sign string) (*UserAuth, error) {
	if !bson.IsObjectIdHex(apiKey) {
		return nil, fmt.Errorf("invalid apiKey.")
	}

	ts, err := strconv.Atoi(timestamp)
	if err != nil {
		return nil, fmt.Errorf("invalid timestamp. %s", err)
	}

	// 如果客户端时间戳小于服务端时间戳5分钟，返错
	if time.Now().After(time.Unix(int64(ts), 0).Add(5 * time.Minute)) {
		log.WithFields(log.Fields{
			"client": timestamp,
			"server": time.Now().Unix(),
		}).Debug("timestamp not match, client too early")
		return nil, fmt.Errorf("invalid timestamp, not match")
	}

	// 如果客户端时间戳大于服务端时间戳1分钟，返错
	if !time.Unix(int64(ts), 0).Add(1 * time.Minute).After(time.Now()) {
		log.WithFields(log.Fields{
			"client": timestamp,
			"server": time.Now().Unix(),
		}).Debug("timestamp not match, client too late")
		return nil, fmt.Errorf("invalid timestamp, not match")
	}

	ua := &UserAuth{}
	err = db.C(global.ColUserAuth).FindId(bson.ObjectIdHex(apiKey)).One(ua)
	if err != nil {
		return nil, err
	}

	// 如果上次调用的时间戳在此次之后，报错
	if ua.Lasted.After(time.Unix(int64(ts), 0)) {
		log.WithFields(log.Fields{
			"latest": ua.Lasted.Unix(),
			"client": ts,
		}).Debug("too early. client again")
		return ua, fmt.Errorf("invalid timestamp")
	}

	// token已经过期
	if ua.Expired.Before(time.Now()) {
		if err := ua.Remove(db); err != nil {
			log.Errorf("remove user-auth %+v failed, %s", ua, err)
			return ua, err
		}
		return ua, fmt.Errorf("Token is expired.")
	}

	// 验证签名的正确性
	hm := hmac.New(sha256.New, []byte(ua.Token))
	r := strings.NewReader(strings.Join([]string{apiKey, timestamp}, ":"))
	io.Copy(hm, r)
	hmRet := fmt.Sprintf("%x", hm.Sum(nil))
	if hmRet != strings.ToLower(sign) {
		log.WithFields(log.Fields{
			"server": hmRet,
			"client": sign,
		}).Debug("not match Signature")
		return nil, fmt.Errorf("not match Signature")
	}

	return ua, nil
}

func (u *UserAuth) Insert(db *mgo.Database) error {
	return db.C(global.ColUserAuth).Insert(u)
}

func (u *UserAuth) Remove(db *mgo.Database) error {
	return db.C(global.ColUserAuth).RemoveId(u.Id)
}

func (u *UserAuth) Update(db *mgo.Database) error {
	update := bson.M{
		"$set": bson.M{
			"lasted": time.Now(),
		},
	}
	return db.C(global.ColUserAuth).UpdateId(u.Id, update)
}
