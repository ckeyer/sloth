package lib

import (
	"fmt"
	"gopkg.in/redis.v3"
)

var redis_cli *redis.Client

func NewClient(host, port string) (client *redis.Client, err error) {
	addr := fmt.Sprintf("%s:%s", host, port)
	opt := &redis.Options{Addr: addr}
	client = redis.NewClient(opt)

	ping := client.Ping()
	if ping.Err() != nil {
		err = fmt.Errorf("faild to connect to  redis(%s), reason:%v", addr, ping.Err())
		return
	}
	redis_cli = client
	return
}
