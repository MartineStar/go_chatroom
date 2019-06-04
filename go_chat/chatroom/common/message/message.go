package message

const(
	LoginMesType			= "LoginMes"
	LoginResMesType			= "LoginResMes"
	RegisterMesType			= "RegisterMes"
	RegisterResMesType 		= "RegisterResMes"
	NotifyUserStatusMesType ="NotifyUserStatusMes"
	SmsMesType 				="SmsMes"
)

//用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type RegisterMes struct{
	User User	`json:"user"`//类型是User结构体
}

type RegisterResMes struct{
	Code int		`json:"code"`	//返回状态码，400表示该用户被占用，200表示注册成功
	//Users []int切片用于将在线用户id返回，用户登录后可以显示在线用户列表
	Users []int

	Error string 	`json:"error"`	//返回错误信息
}

type Message struct{
	Type string		`json:"type"`//消息类型
	Data string		`json:"data"`//消息的类型
}

//定义两个消息，后面需要在增加
type LoginMes struct{
	UserId int	`json:"userId"`
	UserPwd string	`json:"userPwd"`
	UserName string	`json:"userName"`
}


//
type LoginResMes struct{
	Code int   `json:"code"`//返回状态码，500表示该用户未注册，200表示登录成功
	UsersId []int		//保存用户id的切片
	Error string  `json:"error"`//返回的错误信息
}

//服务器推送用户状态变化的消息
type NotifyUserStatusMes struct{
	UserId int 		`json:"userId"`	//用户id
	Status int		`json:"status"`	//用户状态
}

//发送的消息
type SmsMes struct{
	Content string `json:"content"`
	User //匿名结构体
}
