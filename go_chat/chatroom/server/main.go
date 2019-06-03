package main

import (
	"fmt"
	"net"
	"go_chat/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	_ "errors"
	"io"
)

func readPkg(conn net.Conn) (mes message.Message,err error){
	buf := make([]byte,8096)
	fmt.Println("等待读取客户端发送的消息。。。")
	//阻塞等待客户端发送消息，如果客户端断开链接，不再阻塞
	_,err = conn.Read(buf[:4])
	if err != nil{
		// fmt.Println("conn.Read err=",err)
		// err = errors.New("read pkg header error")
		return
	}

	//根据buf[:4]转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	fmt.Println("pkgLen=",pkgLen)
	//根据pkgLen读取消息内容
	_,err = conn.Read(buf[:pkgLen])
	// fmt.Println("n=",n)
	fmt.Println("读取到的buf：",buf[:4])
	if err != nil {
		fmt.Println("conn.Read fail err=,n=",err)
		// err = errors.New("read pkg body error")
		return
	}

	//把pkglen反序列化成message.Message
	err = json.Unmarshal(buf[:pkgLen],&mes)
	
	if err != nil{
		fmt.Println("json.Unmarshal fail,err=",err)
	}

	return
}

func writePkg(conn net.Conn,data []byte) (err error){
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	fmt.Println("before:",buf)

	//此处有疑问
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	
	//发送长度，此处有疑问
	_, err = conn.Write(buf[:4])
	if err != nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	// fmt.Println("after:",buf)

	// fmt.Println("客户端发送消息长度成功,长度为%d,内容为%s",len(data),string(data))

	//发送data本身
	n, err := conn.Write(data)
	if n != int(pkgLen) || err != nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	return
}

//serverProcessLogin，专门处理登录请求
func serverProcessLogin(conn net.Conn,mes *message.Message) (err error){
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
	err =writePkg(conn,data)
	return
}

//serverProcessMes函数，根据客户端发送消息种类的不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn,mes *message.Message) (err error){
	switch mes.Type{
		case message.LoginMesType:
			err = serverProcessLogin(conn,mes)
		case message.RegisterMesType:
			fmt.Println("处理注册")
		default:
			fmt.Println("消息类型不存在，无法处理")
	}
	return
}

//处理和客户段通讯
func process(conn net.Conn){
	//读取客户的消息
	defer conn.Close()
	
	//循环读取客户端发送的消息
	for {
		
		mes,err := readPkg(conn)
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出，服务器端也退出...")
				return
			}else{
				fmt.Println("readPkg err=",err)
			}
		}
		// fmt.Println("mes=",mes)
		err = serverProcessMes(conn,&mes)
		if err != nil{
			return
		}
	}
}

func main(){
	fmt.Println("服务器在8889端口监听....")
	listen,err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil{
		fmt.Println("net.Listen err=",err)
		return
	}

	//一旦监听成功，等待客户端来链接服务器
	for {
		fmt.Println("等待客户端链接服务器...")
		conn,err := listen.Accept()	//conn是一个引用类型
		if err != nil{
			fmt.Println("listen.Accept err=",err)
		}

		//一旦链接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}