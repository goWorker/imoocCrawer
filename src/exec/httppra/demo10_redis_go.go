package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:26679")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	err = conn.Send("auth", "Embe1mpls")
	if err != nil {
		fmt.Println("Wrong password: ", err)
	}
	defer conn.Close()
	_, err = conn.Do("SET", "name", "hanru")
	if err != nil {
		fmt.Println("redis set error:", err)
	}
	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Got name: %s \n", name)
	}
}
