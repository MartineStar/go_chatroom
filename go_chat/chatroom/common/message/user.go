package message

//定义用户的结构体
type User struct {
	//为了序列化和反序列化成功，必须保证用户信息的json字符串可key和结构体字段对应的tag一致
	UserId int	 	`json:"userId"`
	UserPwd string 	`json:"userPwd"`
	UserName string `json:"userName"`
}