package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"	//引入redis包
)

func main(){
	//通过go向redis写入数据和读取数据
	//连接到redis
	conn , err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil{
		fmt.Println("链接错误，err=",err)
	}
	fmt.Println("conn success,conn=",conn)

	// //通过go向redis写入数据,r表示命令执行的结果
	// r,err := conn.Do("set1","name","muzukix")
	// if err != nil{
	// 	fmt.Println("set err=",err,"r =",r)
	// 	return
	// }
	// fmt.Println("set success,r=",r)

	//通过go从redis中读数据
	r,err := redis.String(conn.Do("get","name"))
	if err != nil{
		fmt.Println("get failed,err=",err)
		return
	}
	fmt.Println("get success,r=",r)
}