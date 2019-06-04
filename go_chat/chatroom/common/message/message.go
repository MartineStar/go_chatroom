package message

const(
	LoginMesType		= "LoginMes"
	LoginResMesType		= "LoginResMes"
	RegisterMesType		= "RegisterMes"
	RegisterResMesType 	= "RegisterResMes"
)

type RegisterMes struct{
	User User	`json:"user"`//类型是User结构体
}

type RegisterResMes struct{
	Code int		`json:"code"`	//返回状态码，400表示该用户被占用，200表示注册成功
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
	Error string  `json:"error"`//返回的错误信息
}