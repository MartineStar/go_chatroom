package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//
var pool *redis.Pool
//
func init(){
	pool = &redis.Pool{
		MaxIdle:8,
		MaxActive:0,
		IdleTimeout:100,
		Dial:func()(redis.Conn,error){
			return redis.Dial("tcp","localhost:6379")
		},
	}
}

func main(){
	// pool.Close()
	conn := pool.Get()
	defer conn.Close()
	_,err := conn.Do("set","name","tomcatfklsdjklfd")
	if err != nil{
		fmt.Println("conn.do err = ",err)
	}

	r,err := redis.String(conn.Do("get","name"))
	if err != nil{
		fmt.Println("get failed,err=",err)
		return
	}
	fmt.Println("get success,r=",r)
}