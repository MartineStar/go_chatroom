package process

import (
	"fmt"
	"net"
	"go_chat/chatroom/common/message"
	"go_chat/chatroom/server/util"
	"encoding/json"
)

type UserProcess struct{
	Conn net.Conn
}

//serverProcessLogin，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error){
	//去除mes.Data,并反序列化
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err =",err)
		return
	}


	//1.声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.在声明一个LoginResMes，完成赋值
	// var loginResMes message.LoginResMesType
	var loginResMes message.LoginResMes

	//如果用户id=100，密码为123456，则合法
	if loginMes.UserId ==100 && loginMes.UserPwd =="123456" {
		loginResMes.Code =200
	}else {
		loginResMes.Code =500
		loginResMes.Error = "该用户不存在，请注册在使用"
	}

	//3.将loginResMes序列化
	data,err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("json.Marshal failed,err=",err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes进行序列化，准备发送
	data,err = json.Marshal(resMes)
	if err !=nil{
		fmt.Println("json.Marshal fail,err=",err)
		return
	}

	//6.发送data
	//分层模式mvc，先创建Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err =tf.WritePkg(data)
	return
}