package api

import (
	"github.com/gin-gonic/gin"
	"memorial01/service"
	"net/http"
)

//UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	//申请一个 UserRegister 用户注册服务对象。
	var userRegister service.UserService
	/*
		上下文绑定数据
		ShouldBind能够基于请求的不同，自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象。
		这里将提取的数据绑定到userRegister结构体里面
	*/
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register() //调用这个服务的register方法
		c.JSON(http.StatusOK, res)     //返回这个服务的处理结果
	} else {
		c.JSON(400, err)
	}
}

//UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	//申请一个 UserRegister 用户登录服务对象。
	var userLogin service.UserService

	/*
		上下文绑定数据
		ShouldBind能够基于请求的不同，自动提取JSON、form表单和QueryString类型的数据，并把值绑定到指定的结构体对象。
		这里将提取的数据绑定到userRegister结构体里面
	*/
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()   //调用这个服务的Login方法
		c.JSON(http.StatusOK, res) //返回这个服务的处理结果
	} else {
		c.JSON(400, err)
	}
}
