package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"memorial01/pkg/utils"
	"memorial01/service"
	"net/http"
)

//备忘录接口

//CreateTask 创建备忘录的接口
func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//使用ShouldBind将从POST获取的数据绑定到createTask上
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id) //调用创建方法
		c.JSON(http.StatusOK, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

//ShowTask 展示一条备忘录的详细信息
func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	//使用ShouldBind将从Get获取的数据绑定到showTask上
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("id")) //调用创建方法,传入的是前端路由的id
		c.JSON(http.StatusOK, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

//ListTask 展示所有的备忘录
func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//使用ShouldBind将从POST获取的数据绑定到createTask上
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id) //调用创建方法
		c.JSON(http.StatusOK, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

//UpdateTask 更新备忘录
func UpdateTask(c *gin.Context) {
	var updateask service.UpdateTaskService
	//使用ShouldBind将从Get获取的数据绑定到showTask上
	if err := c.ShouldBind(&updateask); err == nil {
		res := updateask.Update(c.Param("id")) //调用创建方法,传入的是前端路由的id
		c.JSON(http.StatusOK, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

//SearchTask 查找备忘录
func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//使用ShouldBind将从Get获取的数据绑定到showTask上
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claim.Id) //调用创建方法,传入的是前端路由的id
		c.JSON(http.StatusOK, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}

//DeleteTask 删除备忘录
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	//使用ShouldBind将从Get获取的数据绑定到showTask上
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(claim.Id) //调用创建方法,传入的是前端路由的id
		c.JSON(http.StatusOK, res)
	} else {
		logging.Error(err)
		c.JSON(400, err)
	}
}
