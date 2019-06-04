package process

import (
	"fmt"
	"go_chat/chatroom/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)

//在客户端显示当前在线的用户
func outputOnlineUser(){
	fmt.Println("当前在线用户列表:")
	for id,user := range onlineUsers{
		fmt.Println("\t用户id:\t",id)
	}
}

//处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes){
	//先判断是否有这个id的user
	user,ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			userId:notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}