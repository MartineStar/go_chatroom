package main

import (
	"fmt"
	"go_chat/chatroom/client/process"
)

func main(){

	var userId int
	var userPwd string
	//接受用户选择
	var key int

	//判断是否继续显示菜单
	var loop = true

	for loop {
		fmt.Println("------------------欢迎登录多人聊天系统----------------")
		fmt.Println("------------------1. 登录聊天室----------------")
		fmt.Println("------------------2. 注册用户----------------")
		fmt.Println("------------------3. 退出系统----------------")
		fmt.Println("------------------4. 请选择（1-3）----------------")
		//这里必须要换行，不然会出现很多问题
		fmt.Scanf("%d\n",&key)
		switch key{
			case 1:
				fmt.Println("登录聊天室")
				fmt.Println("请输入用户id")
				fmt.Scanf("%d\n",&userId)
				fmt.Println("请输入用户密码")
				fmt.Scanf("%s\n",&userPwd)
				//完成登录
				up := &process.UserProcess{}
				up.Login(userId,userPwd)

				// login(userId,userPwd)
		
				loop = false
			case 2:
				fmt.Println("注册用户")
				loop = false
			case 3:
				fmt.Println("退出系统")
				loop = false
			case 4:
				fmt.Println("你的输入有误，请重新输入")
		}
	}

}