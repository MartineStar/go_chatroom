package process

import (
	"fmt"
	"go_chat/chatroom/common/message"
	"go_chat/chatroom/client/model"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)

//用户登陆成功后，完成CurUser的初始化
var CurUser model.CurUser

//在客户端显示当前在线的用户
func outputOnlineUser(){
	fmt.Println("当前在线用户列表:")
	for id,_ := range onlineUsers{
		fmt.Println("\t用户id:\t",id)
	}
}

//处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes){
	//先判断是否有这个id的user
	user,ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId:notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}