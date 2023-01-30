package main

import (
	"github.com/garyburd/redigo/redis"
)

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	return c
}
