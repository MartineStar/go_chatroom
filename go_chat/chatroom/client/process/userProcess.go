package process

import (
	"fmt"
	"encoding/json"
	// "encoding/binary"
	"net"
	"go_chat/chatroom/common/message"
	"go_chat/chatroom/client/utils"
	"os"
)

type UserProcess struct{

}

func (this *UserProcess) Register(userId int,userPwd string,userName string) (err error){
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
	mes.Type = message.RegisterMesType

	//3.创建一个RegisterMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName=userName

	//4.将registerMes序列化

	data,err := json.Marshal(registerMes)
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
	//创建Transfer实例
	tf := &utils.Transfer{
		Conn :conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("register writePkg err=",err)
	}

	//接收服务端的消息
	mes,err = tf.ReadPkg()
	
	if err != nil{
		fmt.Println("register readPkg(conn) err =",err)
		return
	}

	//将mes的Data部分反序列化为registerResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code ==200{
		fmt.Println("注册成功，请重新登录！")
		os.Exit(0)
	}else{
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return

}

//写一个函数，完成登录,返回类型为error，能充分描述错误，如果返回bool值的话就比较片面
func (this *UserProcess) Login(userId int,userPwd string) (err error){
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
	//创建Transfer实例
	tf := &utils.Transfer{
		Conn :conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("writePkg err=",err)
	}

	//接收服务端的消息
	mes,err = tf.ReadPkg()
	
	if err != nil{
		fmt.Println("readPkg(conn) err =",err)
		return
	}

	//将mes的Data部分反序列化为LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code ==200{
		fmt.Println("登陆成功")

		//显示当前在线用户列表
		fmt.Println("当前在线用户列表如下:")
		for _,v := range loginResMes.UsersId{
			//跳过自己
			if v == userId{
				continue
			}
			fmt.Println("\t用户id:\t",v)

			//完成客户端的onlineUsers的初始化
			user := &message.User{
				UserId : v,
				UserStatus:message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("\n\n")
		//客户端启动一个协程，该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接受并显示在客户端的终端
		//1.循环显示菜单
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	}else{
		fmt.Println(loginResMes.Error)
	}
	return 
}

