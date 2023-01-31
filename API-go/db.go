package main

import (
	"context"
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

func postRedis(ctx context.Context, funcionario Funcionario) {

	select {
	case <-ctx.Done():
		fmt.Println("Sinal cancelamento")
		return
	default:
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

}

func getAllRedis(ctx context.Context) []Funcionario {

	select {
	case <-ctx.Done():
		fmt.Println("Sinal cancelamento")
		return nil
	default:
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
}

func findRedis(ctx context.Context, id string) Funcionario {

	select {
	case <-ctx.Done():
		var funcionario Funcionario
		fmt.Println("Sinal cancelamento")
		return funcionario
	default:
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

}

func deleteRedis(ctx context.Context, id string) Funcionario {

	select {
	case <-ctx.Done():
		var funcionario Funcionario
		fmt.Println("Sinal cancelamento")
		return funcionario
	default:
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
		replyDel, err := c.Do("DEL", "post:"+id)
		if err != nil {
			panic(err)
		}
		fmt.Println(replyDel)
		return funcionario
	}
}
