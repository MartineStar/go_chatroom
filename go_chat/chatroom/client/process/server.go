package process

import (
	"fmt"
	"os"
	"net"
	"go_chat/chatroom/client/utils"
	"go_chat/chatroom/common/message"
	"encoding/json"
)

//显示登录成功之后的界面
func ShowMenu(){
	fmt.Println("--------恭喜xxx登录成功---------")
	fmt.Println("--------1.显示在线用户列表-------")
	fmt.Println("--------2.发 送 消 息----------")
	fmt.Println("--------3.信 息 列 表----------")
	fmt.Println("--------4.退 出 系 统----------")
	fmt.Print("----------请选择(1-4):")

	var key int
	fmt.Scanf("%d\n",&key)
	switch key{
		case 1:
			fmt.Println("显示在线用户列表")
			outputOnlineUser()
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("你选择退出系统，系统正在退出...")
			os.Exit(0)
		default:
			fmt.Println("你输入的选项不正确，请重新输入")
	}

}

func serverProcessMes(conn net.Conn){
	tf := &utils.Transfer{Conn:conn,}
	//创建transfer实例，不停地读取服务端的发送的消息
	for {
		fmt.Printf("客户端正在读取服务器发送的消息...")
		mes,err := tf.ReadPkg()
		if err != nil{
			fmt.Println("tf.ReadPkg err=",err)
			return
		}

		//如果读取到消息，进行下一步
		switch mes.Type {
			case message.NotifyUserStatusMesType:
				//1.取出NotifyUserStatusMes
				var notifyUserStatusMes message.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
				//2.把这个用户的信息，状态保存到客户map[int]User 中
				updateUserStatus(&notifyUserStatusMes)

			default:
				fmt.Println("服务器返回了未知的消息类型")
		}
		fmt.Printf("mes=%v\n",mes)
	}
}