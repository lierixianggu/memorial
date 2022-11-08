package models

func migrate() {
	//自动迁移模式
	/*
		AutoMigrate函数是把代码映射到数据库中
		AddForeignKey函数是添加外键
	*/
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
}
