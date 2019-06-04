package main

import (
	"fmt"
	"net"
	"go_chat/chatroom/common/message"
	"go_chat/chatroom/server/process"
	"go_chat/chatroom/server/util"
	"io"
)

//创建Processor结构体
type Processor struct{
	Conn net.Conn
}

//serverProcessMes函数，根据客户端发送消息种类的不同，决定调用哪个函数来处理
func (this *Processor)serverProcessMes(mes *message.Message) (err error){
	fmt.Println(mes)
	switch mes.Type{
		case message.LoginMesType:
			//创建和一个userProcess实例，处理登录
			up := &process.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)

		case message.RegisterMesType:
			up := &process.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessRegister(mes)

		case message.SmsMesType:
			smsProcess := &process.SmsProcess{}
			smsProcess.SendGroupMes(mes)
		default:
			fmt.Println("消息类型不存在，无法处理")
	}
	return
}


func (this *Processor) process2() (err error){
	//循环读取客户端发送的消息
	for {
		tf := &utils.Transfer{
			Conn:this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出，服务器端也退出...")
				//这里要指明返回err，否则报错err is shadowed during return
				// return
				return err
			}else{
				fmt.Println("readPkg err=",err)
			}
		}
		// fmt.Println("mes=",mes)
		err = this.serverProcessMes(&mes)
		if err != nil{
			return err
		}
	}
	
}