package utils

// import (
// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// type Config struct {
// 	Host     string
// 	Port     string
// 	UserName string
// 	Password string
// 	Database string
// }

// // mongo的辅助类
// type Mdb struct {
// 	*Config
// 	baseSession *mgo.Session
// }

// func NewMdbWithHost(Host string) *Mdb {
// 	return NewMdb(Host, "", "", "", "")
// }

// func NewMdb(host, port, database, userName, password string) *Mdb {
// 	mdb := &Mdb{
// 		Config: &Config{
// 			Host:     host,
// 			Port:     port,
// 			UserName: userName,
// 			Password: password,
// 			Database: database,
// 		},
// 	}
// 	mdb.connect()
// 	return mdb
// }

// func NewMdbWithConf(c *Config) (mdb *Mdb) {
// 	mdb = &Mdb{
// 		Config: c,
// 	}
// 	mdb.connect()
// 	return
// }

// func (m *Mdb) connect() {
// 	// 连接url ： [mongodb:// ][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
// 	url := m.Host
// 	if m.UserName != "" && m.Password != "" {
// 		url = m.UserName + ":" + m.Password + "@" + url
// 	}
// 	if m.Port != "" {
// 		url = url + ":" + m.Port
// 	}
// 	if m.Database != "" {
// 		url = url + "/" + m.Database
// 	}
// 	var err error
// 	m.baseSession, err = mgo.Dial(url)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func (m *Mdb) Session() *mgo.Session {
// 	return m.baseSession.New()
// }

// func (m *Mdb) DB(s *mgo.Session) *mgo.Database {
// 	return s.DB(m.Config.Database)
// }

// func (m *Mdb) WithC(collection string, job func(*mgo.Collection) error) error {
// 	s := m.baseSession.New()
// 	defer s.Close()
// 	return job(s.DB(m.Config.Database).C(collection))
// }

// func (m *Mdb) Upsert(collection string, selector interface{}, change interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		_, err := c.Upsert(selector, change)
// 		return err
// 	})
// }

// func (m *Mdb) UpdateId(collection string, id interface{}, change interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		return c.UpdateId(id, change)
// 	})
// }
// func (m *Mdb) Update(collection string, selector, change interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		return c.Update(selector, change)
// 	})
// }
// func (m *Mdb) UpdateAll(collection string, selector, change interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		_, err := c.UpdateAll(selector, change)
// 		return err
// 	})
// }

// func (m *Mdb) Insert(collection string, data ...interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		return c.Insert(data...)
// 	})
// }

// func (m *Mdb) All(collection string, query interface{}, result interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		return c.Find(query).All(result)
// 	})
// }

// // 返回所有复合 query 条件的item， 并且被 projection 限制返回的fields
// func (m *Mdb) AllSelect(collection string, query interface{}, projection interface{}, result interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		return c.Find(query).Select(projection).All(result)
// 	})
// }

// func (m *Mdb) One(collection string, query interface{}, result interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		return c.Find(query).One(result)
// 	})
// }

// func (m *Mdb) OneSelect(collection string, query interface{}, projection interface{}, result interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		return c.Find(query).Select(projection).One(result)
// 	})
// }

// // 等效于: m.One(collection,bson.M{"_id":id},result)
// func (m *Mdb) FindId(collection string, id interface{}, result interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		return c.Find(bson.M{"_id": id}).One(result)
// 	})
// }

// func (m *Mdb) RemoveId(collection string, id interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		err := c.RemoveId(id)
// 		return err
// 	})
// }
// func (m *Mdb) Remove(collection string, selector interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		err := c.Remove(selector)
// 		return err
// 	})
// }
// func (m *Mdb) RemoveAll(collection string, selector interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		_, err := c.RemoveAll(selector)
// 		return err
// 	})
// }

// func (m *Mdb) CountId(collection string, id interface{}) (n int) {
// 	m.WithC(collection, func(c *mgo.Collection) error {
// 		var err error
// 		n, err = c.FindId(id).Count()
// 		return err
// 	})
// 	return n
// }
// func (m *Mdb) Count(collection string, query interface{}) (n int) {
// 	m.WithC(collection, func(c *mgo.Collection) error {
// 		var err error
// 		n, err = c.Find(query).Count()
// 		return err
// 	})
// 	return n
// }
// func (m *Mdb) Exist(collection string, query interface{}) bool {
// 	return m.Count(collection, query) != 0
// }
// func (m *Mdb) ExistId(collection string, id interface{}) bool {
// 	return m.CountId(collection, id) != 0
// }

// func (m *Mdb) Page(collection string, query bson.M, offset int, limit int, result interface{}) error {
// 	return m.WithC(collection, func(c *mgo.Collection) error {
// 		return c.Find(query).Skip(offset).Limit(limit).All(result)
// 	})
// }

// // 获取页面数据和“所有”符合条件的记录“总共”的条数
// func (m *Mdb) PageAndCount(collection string, query bson.M, offset int, limit int, result interface{}) (total int, err error) {
// 	err = m.WithC(collection, func(c *mgo.Collection) error {
// 		total, err = c.Find(query).Count()
// 		if err != nil {
// 			return err
// 		}
// 		return c.Find(query).Skip(offset).Limit(limit).All(result)
// 	})
// 	return total, err
// }

// // 等同与UpdateId(collection,id,bson.M{"$set":change})
// func (m *Mdb) SetId(collection string, id interface{}, change interface{}) error {
// 	return m.UpdateId(collection, id, bson.M{"$set": change})
// }
