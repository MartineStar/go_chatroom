package model


import (
	"net"
	"go_chat/chatroom/common/message"
)

//客户端很多地方会使用curUser，设置为全局

type CurUser struct{
	Conn net.Conn
	message.User
}