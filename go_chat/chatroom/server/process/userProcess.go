package process

import (
	"fmt"
	"net"
	"go_chat/chatroom/common/message"
	"go_chat/chatroom/server/util"
	"encoding/json"
	"go_chat/chatroom/server/model"
)

type UserProcess struct{
	Conn net.Conn
	UserId int	//旨在指明这个conn是属于哪个用户
}

//单独通知的方法
func (this *UserProcess) NotifyMeOnline(userId int){
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//序列化
	data,err := json.Marshal(notifyUserStatusMes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	mes.Data = string(data)

	//序列化mes
	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	//发送
	tf := &utils.Transfer{
		Conn :this.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("notyfymeOnline err=",err)
		return
	}

}


//通知所有在线用户的方法,userId要通知其他用户我上线了
func (this *UserProcess) NotifyOthersOnlineUser(userId int){
	//遍历onlineUsers,然后一个个发送NotifyUserStatusMes
	for id,up := range userMgr.onlineUsers{
		if id == userId{
			continue
		}
		//开始通知
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error){
	//取出mes.Data,并反序列化
	var registerMes message.RegisterMes
	json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil{
		fmt.Println("register json.Unmarshal fail err =",err)
		return
	}

	//1.声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes

	//2.去redis中校验并存入用户信息，完成注册
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil{
		if err == model.ERROR_USER_EXIST{
			registerResMes.Code = 505
			//Error()方法可以取出error中的错误描述
			registerResMes.Error = model.ERROR_USER_EXIST.Error()
		}else{
			registerResMes.Code = 506
			registerResMes.Error = "注册时发生未知错误"
		}
	}else{
		registerResMes.Code = 200	

	}

	//3.将registerResMes序列化
	data,err := json.Marshal(registerResMes)
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


//serverProcessLogin，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error){
	//取出mes.Data,并反序列化
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil{
		fmt.Println("login json.Unmarshal fail err =",err)
		return
	}


	//1.声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType



	//2.在声明一个LoginResMes，完成赋值
	// var loginResMes message.LoginResMesType
	var loginResMes message.LoginResMes

	//2.去redis中核对用户信息
	user,err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	
	if err != nil{
		if err == model.ERROR_USER_NOTEXISTS{
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_NOTEXISTS{
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		}else{
			loginResMes.Code = 500
			loginResMes.Error = "该用户不存在，请注册再使用..."
		}

	}else{
		loginResMes.Code = 200
		//将登录成功的用户的userId放入对应的UserProcess实例
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)

		//通知其他用户我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)

		//遍历userMgr.onlineUsers,获取userId
		for id,_ := range userMgr.onlineUsers{
			loginResMes.UsersId = append(loginResMes.UsersId,id)
		}
		

		fmt.Println(user,"登录成功")
	}
	// //如果用户id=100，密码为123456，则合法
	// if loginMes.UserId ==100 && loginMes.UserPwd =="123456" {
	// 	loginResMes.Code =200
	// }else {
	// 	loginResMes.Code =500
	// 	loginResMes.Error = "该用户不存在，请注册在使用"
	// }

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