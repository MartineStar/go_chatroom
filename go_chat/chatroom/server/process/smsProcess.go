package process

import (
	"fmt"
	"go_chat/chatroom/common/message"
	"encoding/json"
	"net"
	"go_chat/chatroom/server/util"
)

type SmsProcess struct{}

//转发消息
func (this *SmsProcess) SendGroupMes(mes *message.Message){

	//取出mes的内容SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	data,err := json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	//遍历服务器端的onlineUsers map[int]*UserProcess
	for id,up := range userMgr.onlineUsers{
		//跳过自己
		if id == smsMes.UserId{
			continue
		}
		this.SendMesToEachOnlineUser(data,up.Conn)


	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte,conn net.Conn){
	tf := &utils.Transfer{
		Conn:conn,
	}

	err := tf.WritePkg(data)
	if err != nil{
		fmt.Println("转发消息失败，err=",err)
	}
}