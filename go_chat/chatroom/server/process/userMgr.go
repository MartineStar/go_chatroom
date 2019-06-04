package process

import (
	"fmt"
)

//UserMgr实例在服务器端有且仅有一个，并在很多地方可以用到，定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct{
	onlineUsers map[int]*UserProcess
}

//完成对userMgr的初始化
func init()  {
	userMgr = &UserMgr{
		onlineUsers:make(map[int]*UserProcess,1024),
	}
}

//完成对onlineUsers添加
func (this *UserMgr) AddOnlineUser(up *UserProcess){
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DelOnlineUser(userId int){
	delete(this.onlineUsers,userId)
}

//返回当前所有在线的用户
func (this *UserMgr) GetAllOnlines() map[int]*UserProcess{
	return this.onlineUsers
}
//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess,err error){
	up,ok := this.onlineUsers[userId]
	if !ok{
		//fmt.Errorf()  格式化输出错误
		err = fmt.Errorf("用户%d 不存在",userId)
		return
	}
	return
}