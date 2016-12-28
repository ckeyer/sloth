package global

import (
	"fmt"

	"github.com/ckeyer/commons/db/mongo"
	"gopkg.in/mgo.v2"
)

const (
	// mgo.Collection.Name
	ColUser     = "user"
	ColProject  = "project"
	ColWebhook  = "webhook"
	ColUserAuth = "user_auth"
	ColSettings = "settings"
)

var (
	// for reconnect
	mgourl = ""
)

func mgoIndexes() mongo.MgoIndexs {
	return mongo.MgoIndexs{
		ColUser: []mgo.Index{
			mgo.Index{
				Key:    []string{"name", "email", "phone"},
				Unique: true,
			},
		},
		ColUserAuth: []mgo.Index{
			mgo.Index{
				Key: []string{"user_id", "lasted", "expired"},
			},
		},
		ColProject: []mgo.Index{
			mgo.Index{
				Key:    []string{"name"},
				Unique: true,
			},
			mgo.Index{
				Key: []string{"type"},
			},
		},
		ColSettings: []mgo.Index{
			mgo.Index{
				Key:    []string{"key"},
				Unique: true,
			},
		},
	}
}

// InitDB
func InitDB(url string) (*mgo.Database, error) {
	db, err := mongo.NewMdbWithURL(url)
	if err != nil {
		return nil, fmt.Errorf("Connect %s failed, %s", url, err)
	}

	indexes := mgoIndexes()
	err = indexes.Setup(db)
	if err != nil {
		return nil, fmt.Errorf("Set MGO Indexes failed, %s", err)
	}

	mgourl = url

	return db, nil
}
