package main

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	return c
}

func postRedis(funcionario Funcionario) {

	c := RedisConnect()
	defer c.Close()
	object, err := json.Marshal(funcionario)
	if err != nil {
		panic(err)
	}
	reply, err := c.Do("SET", "post:"+funcionario.Id, object)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}

func getAllRedis() []Funcionario {

	var funcionarios []Funcionario
	c := RedisConnect()
	defer c.Close()
	keys, err := c.Do("KEYS", "post:*")
	if err != nil {
		panic(err)
	}
	for _, i := range keys.([]interface{}) {

		var funcionario Funcionario
		reply, err := c.Do("GET", i.([]byte))
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(reply.([]byte), &funcionario); err != nil {
			panic(err)
		}
		funcionarios = append(funcionarios, funcionario)
	}
	return funcionarios
}

func findRedis(id string) Funcionario {

	var funcionario Funcionario
	c := RedisConnect()
	defer c.Close()
	reply, err := c.Do("GET", "post:"+id)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(reply.([]byte), &funcionario); err != nil {
		panic(err)
	}
	return funcionario
}

func deleteRedis(id string) Funcionario {

	funcionario := findRedis(id)
	c := RedisConnect()
	defer c.Close()
	_, err := c.Do("DEL", "post:"+id)
	if err != nil {
		panic(err)
	}
	return funcionario
}
