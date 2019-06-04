package process

import (
	"fmt"
	"go_chat/chatroom/common/message"
	"go_chat/chatroom/client/utils"
	"encoding/json"
)

type SmsProcess struct{


}

//发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error){
	// 1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2.创建smsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId=CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化
	data,err :=json.Marshal(smsMes)
	if err != nil{
		fmt.Println("sendGroupMes json.Marshal fail ,err=",err.Error())
		return
	}

	//4.再次序列化
	mes.Data = string(data)

	data,err =json.Marshal(mes)
	if err != nil{
		fmt.Println("sendGroupMes json.Marshal fail ,err=",err.Error())
		return
	}

	//5.发送mes给服务器
	tf := &utils.Transfer{
		Conn:CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("sendGroupMes err=",err.Error())
		return
	}
	return

}