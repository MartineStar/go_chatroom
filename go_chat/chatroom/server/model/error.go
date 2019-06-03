package model

import (
	"errors"
)
//根据业务逻辑的需要，自定义一些错误

var (
	ERROR_USER_NOTEXISTS = 	errors.New("用户不存在")
	ERROR_USER_EXIST = errors.New("该用户名已被占用")
	ERROR_USER_PWD = errors.New("密码错误")
)