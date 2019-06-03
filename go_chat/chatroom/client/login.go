package main

import (
	"fmt"
	"encoding/json"
	// "encoding/binary"
	"net"
	"go_chat/chatroom/common/message"
)

//写一个函数，完成登录,返回类型为error，能充分描述错误，如果返回bool值的话就比较片面
func login(userId int,userPwd string) (err error){
	// //开始定协议
	// fmt.Printf("userId=%d,userPwd=%s\n",userId,userPwd)
	// return nil

	//1.链接到服务器
	conn,err := net.Dial("tcp","localhost:8889")
	if err != nil{
		fmt.Println("net.Dial err=",err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd


	//4.将loginMes序列化

	data,err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json.Marshal 1 err=",err)
		return
	}

	//5.把data赋值给mes.Data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal 2 err=",err)
		return
	}
	//向服务器发送数据
	err = writePkg(conn,data)
	if err != nil{
		fmt.Println("writePkg err=",err)
	}

	//接收服务端的消息
	mes,err = readPkg(conn)
	
	if err != nil{
		fmt.Println("readPkg(conn) err =",err)
		return
	}



	//将mes的Data部分反序列化为LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code ==200{
		fmt.Println("登陆成功")

	}else if loginResMes.Code ==500{
		fmt.Println(loginResMes.Error)
	}
	return 
}