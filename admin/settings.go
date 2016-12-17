package admin

import (
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/ckeyer/sloth/global"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 一个settings缓存
var setsManage struct {
	sync.Mutex
	kv map[string]string
}

func init() {
	setsManage.kv = map[string]string{}
}

type Setting struct {
	Key     string    `json:"key" bson:"key"`
	Value   string    `json:"value" bson:"value"`
	Created time.Time `json:"created" bson:"created"`
}

func SetKV(db *mgo.Database, k, v string) error {
	s := &Setting{
		Key:     k,
		Value:   v,
		Created: time.Now(),
	}
	_, err := db.C(global.ColSettings).Upsert(bson.M{"key": k}, s)
	if err != nil {
		log.Errorf("upsert settings failed, %s", err)
		return err
	}

	setsManage.Lock()
	setsManage.kv[k] = v
	setsManage.Unlock()

	log.WithFields(log.Fields{
		"key":   k,
		"value": v,
	}).Debugf("upsert settings")
	return nil
}

func GetValue(db *mgo.Database, k string) (string, error) {
	setsManage.Lock()
	defer setsManage.Unlock()

	if v, ok := setsManage.kv[k]; ok {
		return v, nil
	}

	s := &Setting{}
	err := db.C(global.ColSettings).Find(bson.M{"key": k}).One(s)
	if err != nil {
		return "", err
	}
	setsManage.kv[k] = s.Value
	return s.Value, nil
}

func GetKVs(db *mgo.Database, ks ...string) (map[string]string, error) {
	query := bson.M{
		"key": bson.M{
			"$in": ks,
		},
	}
	var kvs []Setting
	err := db.C(global.ColSettings).Find(query).All(&kvs)
	if err != nil {
		return nil, err
	}

	ret := map[string]string{}
	setsManage.Lock()
	defer setsManage.Unlock()
	for _, set := range kvs {
		ret[set.Key] = set.Value
		setsManage.kv[set.Key] = set.Value
	}

	return ret, nil
}
