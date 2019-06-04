package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"go_chat/chatroom/common/message"
)


//定义userDao结构体，完成对user的操作
type UserDao struct{
	pool *redis.Pool
}

//在服务器启动后，初始化一个userDa实例，设置为全局变量，在需要和redis操作时，直接使用
var (
	MyUserDao *UserDao
)


//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao){
	userDao = &UserDao{
		pool:pool,
	}
	return userDao
}

//
func (this *UserDao) getUserById(conn redis.Conn,id int) (user *User,err error){
	//通过给定的id去redis查询用户
	res,err := redis.String(conn.Do("HGet","users",id))
	if err != nil{
		//错误
		if err == redis.ErrNil{    //表示在users哈希中，没有对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	//把res反序列化成Users实例
	err = json.Unmarshal([]byte(res),user)
	if err != nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return
}

//完成用户注册
func (this *UserDao) Register(user *message.User) (err error){
	//先从UserDao连接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	_,err = this.getUserById(conn,user.UserId)
	if err == nil{	//如果不报错，则说明这个id存在了
		err = ERROR_USER_NOTEXISTS
		return
	}

	//序列化，入库，完成注册
	data,err := json.Marshal(user)
	if err != nil{
		return
	}

	//入库
	_,err = conn.Do("Hset","users",user.UserId,string(data))
	if err != nil{
		fmt.Println("保存注册用户错误，err=",err)
		return
	}
	return

}

//完成用户登录校验
func (this *UserDao) Login(userId int,userPwd string) (user *User,err error){
	//先从UserDao连接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()
	user,err = this.getUserById(conn,userId)
	if err != nil{
		return
	}
	//至此说明这个用户账号是存在的
	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		return
	}
	return

}