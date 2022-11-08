package service

import (
	"memorial01/models"
	"memorial01/serializers"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做,1是已做
}

type ShowTaskService struct {
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做,1是已做
}

type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type DeleteTaskService struct {
}

//Create 是创建备忘录的方法
func (service *CreateTaskService) Create(id uint) serializers.Response {
	var user models.User
	code := 200
	models.DB.First(&user, id)
	task := models.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,   //从前端获得
		Status:    0,               //默认初始值
		Content:   service.Content, //从前端获得
		StartTime: time.Now().Unix(),
		EndTime:   0, //默认值
	}
	err := models.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializers.Response{
			Status: code,
			Msg:    "创建备忘录失败",
		}
	}
	return serializers.Response{
		Status: code,
		Msg:    "备忘录创建成功",
	}
}

//展示一条备忘录

//Show 是查询一条备忘录详细信息的方法
func (service *ShowTaskService) Show(tid string) serializers.Response {
	var task models.Task
	code := 200
	err := models.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializers.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializers.Response{
		Status: code,
		Data:   serializers.BuildTask(task),
		Msg:    "查询成功",
	}

}

//List 列表返回用户所有备忘录
func (service *ListTaskService) List(uid uint) serializers.Response {
	var tasks []models.Task
	count := 0
	if service.PageSize == 0 { //如果传过来的页数是0
		service.PageSize = 15 //默认每页15个
	}
	models.DB.Model(&models.Task{}).Preload("User").Where("uid=?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	//返回到前端的内容需要序列化,gin.H{}这类的
	return serializers.BuildListResponse(serializers.BuildTasks(tasks), uint(count))
}

//Update 更新备忘录的方法
func (service *UpdateTaskService) Update(tid string) serializers.Response {
	var task models.Task
	models.DB.First(&task, tid)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	models.DB.Save(&task)
	return serializers.Response{
		Status: 200,
		Data:   serializers.BuildTask(task),
		Msg:    "数据更新完毕",
	}
}

//Search 查找备忘录的方法
func (service *SearchTaskService) Search(uid uint) serializers.Response {
	var tasks []models.Task
	count := 0

	if service.PageSize == 0 {
		service.PageSize = 15
	}
	//进行模糊查询
	//这里同时进行了分页操作
	models.DB.Model(&models.Task{}).Preload("User").Where("uid=?", uid). //预加载与外表关联找到用户
										Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").                //找到用户之后进行模糊搜索
										Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks) //然后这里进行计数操作，并且分页
	return serializers.BuildListResponse(serializers.BuildTasks(tasks), uint(count))

}

//Delete 删除备忘录的方法
func (service *DeleteTaskService) Delete(uid uint) serializers.Response {
	var task models.Task
	err := models.DB.Delete(&task, uid).Error
	if err != nil {
		return serializers.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializers.Response{
		Status: 200,
		Msg:    "删除成功",
	}

}
