package initialize

import (
	"fmt"
	"go-gin/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDB() {
	var err error
	var myHost = IniConf.String("my_host")
	var myUser = IniConf.String("my_user")
	var myPass = IniConf.String("my_pass")
	var myDb = IniConf.String("my_db")
	var myOptions = IniConf.String("my_options")
	dsn := myUser + ":" + myPass + "@tcp(" + myHost + ")/" + myDb + "?" + myOptions
	fmt.Println(dsn)
	global.DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		log.Panicln("err:", err.Error())
	} else {
		sqlDB, err1 := global.DB.DB()
		if err1 != nil {
			log.Panicln("err1:", err1.Error())
		}
		var maxIdleConns, _ = IniConf.Int("maxidleconns")
		var maxOpenConns, _ = IniConf.Int("maxopenconns")
		sqlDB.SetMaxOpenConns(maxOpenConns)
		sqlDB.SetMaxIdleConns(maxIdleConns)
	}

}

func CloseDB() {
	sqlDB, _ := global.DB.DB()
	sqlDB.Close()
}
