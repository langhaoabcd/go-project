package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	conf "go-project/configs"
)

var db *gorm.DB

func GetDB() *gorm.DB  {
	return db
}

func init() {
	var err error
	db, err = gorm.Open("mysql", conf.GetConf().Dbconn)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to  database, but got err=%+v", err))
	}
	db.SingularTable(true) //全局禁用表名复数 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	//db.DB().SetConnMaxLifetime(1 * time.Second)
	//db.DB().SetMaxIdleConns(20) //最大连接数
	//db.DB().SetMaxOpenConns(2000)//最大闲置个数
	db.LogMode(true)//显示详细日志
}