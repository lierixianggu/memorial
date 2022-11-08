package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"memorial01/api"
	"memorial01/middleware"
)

func NewRouter() *gin.Engine {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 创建基于 cookie 的存储引擎，something-very-secret 参数是用于加密的密钥
	store := cookie.NewStore([]byte("something-very-secret"))
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，可以替换成其他存储引擎
	r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("api/v1")
	{
		//用户操作(注册/登录)
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/")
		authed.Use(middleware.JWT()) //使用中间件JWT去验证用户有没有权限
		{
			authed.POST("task", api.CreateTask)
			authed.GET("task/:id", api.ShowTask)
			authed.GET("tasks", api.ListTask)
			authed.PUT("task/:id", api.UpdateTask)
			authed.POST("search", api.SearchTask)
			authed.DELETE("task/:id", api.DeleteTask)
		}
	}
	return r
}
