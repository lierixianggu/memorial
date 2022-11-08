package service

import (
	"memorial01/models"
	"memorial01/pkg/utils"
	"memorial01/serializers"
)

//UserService service层编写方法
type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

//Register 用户注册的方法
func (service *UserService) Register() serializers.Response {
	var user models.User
	var count int
	//1.先对传过来的用户名进行验证，查看是否已经存在了
	models.DB.Model(&models.User{}).Where("user_name=?", service.UserName).
		First(&user).Count(&count)
	if count == 1 {
		return serializers.Response{
			Status: 400,
			Msg:    "已经有这个人了，不需要再注册",
		}
	}
	user.UserName = service.UserName

	//对密码进行加密
	err := user.SetPassword(service.Password)
	if err != nil {
		return serializers.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}

	//创建用户
	err = models.DB.Create(&user).Error
	if err != nil {
		return serializers.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return serializers.Response{
		Status: 200,
		Msg:    "用户注册成功！",
	}
}

//Login 用户登陆的方法
func (service *UserService) Login() serializers.Response {
	var user models.User

	//1.先去找一下这个user，看看数据库有没有这个用户
	err := models.DB.Where("user_name=?", service.UserName).First(&user).Error
	if err != nil {
		return serializers.Response{
			Status: 400,
			Msg:    "用户不存在，请先注册",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return serializers.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//发一个token，为了其他功能需要身份验证所给前端存储的
	//创建一个备忘录,这个功能需要token，不然都不知道是谁创建的备忘录
	//我们传入用户的id，username和password以及authority权限加密成token，后续就可以进行身份的验证
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serializers.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializers.Response{
		Status: 200,
		Data: serializers.TokenData{ //带有token的返回值
			User:  serializers.BuildUser(user),
			Token: token,
		},
		Msg: "登录成功",
	}
}
