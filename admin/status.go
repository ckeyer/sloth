package admin

import (
	"github.com/ckeyer/sloth/global"
	"gopkg.in/mgo.v2"
)

func Status(db *mgo.Database) (map[string]interface{}, error) {
	users, err := db.C(global.ColUser).Find(nil).Count()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"app":     "sloth",
		"user":    users,
		"project": 123,
	}, nil
}
