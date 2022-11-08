package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

/*
	通过conf.go获得的路径,进行数据库的连接
*/

var DB *gorm.DB

//接收config.go中的path
//进行数据库连接

func Database(connstring string) {
	fmt.Println("connstring:", connstring)
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Println(err)
		fmt.Println("数据库连接错误")
		return
	}
	fmt.Println("数据库连接成功")
	db.SingularTable(true)      //表名字不加s
	db.DB().SetMaxIdleConns(20) //设置连接池
	db.DB().SetMaxOpenConns(10) //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migrate()
}
