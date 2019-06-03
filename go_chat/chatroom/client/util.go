package main
import (
	"fmt"
	"net"
	"go_chat/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
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