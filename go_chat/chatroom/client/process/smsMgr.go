package process

import (
	"fmt"
	"go_chat/chatroom/common/message"
	"encoding/json"

)

func outputGroupMes(mes *message.Message){
	//反序列化mes.Data 
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("json.Unmarshal err=",err.Error())
		return
	}

	//显示信息
	info := fmt.Sprintf("用户[%d]说: %s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}