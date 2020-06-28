package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"

	//"database/sql"
	"github.com/jinzhu/gorm"
)

var (
	//Db  *sql.DB
	DbHelper  *gorm.DB
	err error
)

//func init() {
//	Db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/bookstore0612")
//	if err != nil {
//		panic(err.Error())
//	}
//}

//func init() {
//	Db, err = sql.Open("mysql", "root:Dev@8888@tcp(192.168.50.100:3306)/game_center")
//	if err != nil {
//		panic(err.Error())
//	}
//}

func init() {
	DbHelper, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/game_center?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	// 最大连接数
	DbHelper.DB().SetMaxOpenConns(100)
	// 闲置连接数
	DbHelper.DB().SetMaxIdleConns(20)
	// 最大连接周期
	DbHelper.DB().SetConnMaxLifetime(100 * time.Second)

	DbHelper.LogMode(true) //打开日志
	DbHelper.SingularTable(true) //框架自带生成的对象是表名的英文复数形式 这个属性是可以不让带复数
}
