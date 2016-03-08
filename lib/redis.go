package lib

import (
	"fmt"
	"gopkg.in/redis.v3"
)

var (
	redisOption *redis.Options
	redisCli    *redis.Client
)

func InitRedis(host, port string) {
	addr := fmt.Sprintf("%s:%s", host, port)
	redisOption = &redis.Options{
		Network: "tcp",
		Addr:    addr,
	}
	_, err := NewClient(redisOption)
	if err != nil {
		log.Fatal("redis init error, ", err)
		return
	}
}

func NewClient(opt *redis.Options) (client *redis.Client, err error) {
	client = redis.NewClient(opt)

	ping := client.Ping()
	if ping.Err() != nil {
		err = fmt.Errorf("faild to connect to  redis(%s), reason:%v", opt.Addr, ping.Err())
		return
	}
	redisCli = client
	return
}
