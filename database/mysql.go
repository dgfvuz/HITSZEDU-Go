package database

import (
	"fmt"
	"hitszedu-go/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectMysql() {
	_db, err := gorm.Open(config.GetString("database.type"),
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			config.GetString("database.username"),
			config.GetString("database.password"),
			config.GetString("database.address"),
			config.GetString("database.port"),
			config.GetString("database.dbname"),
			config.GetString("database.charset")))

	if err != nil {
		panic(err)
	}
	_db.SingularTable(true)
	_db.DB().SetMaxIdleConns(10)
	_db.DB().SetMaxOpenConns(100)
	_db.LogMode(config.GetBool("database.logmode"))
	DB = _db
	fmt.Println("数据库连接成功!")
}

// 封装了db.Create, 返回error
func Create(value interface{}) error {
	return DB.Create(value).Error
}

// 封装了db.Where.First, 返回error
func First(out interface{}, query interface{}, args ...interface{}) error {
	return DB.Where(query, args...).First(out).Error
}

// 封装了db.Where.Find, 返回error
func Find(out interface{}, query interface{}, args ...interface{}) error {
	return DB.Where(query, args...).Find(out).Error
}

// 封装了db. Save, 返回error
func Save(value interface{}) error {
	return DB.Save(value).Error
}

//  封装了db.Update, 返回error
func Update(value interface{}, attrs ...interface{}) error {
	return DB.Model(value).Update(attrs...).Error
}

//  封装了db.Updates, 返回error
func Updates(entity interface{}, values interface{}) error {
	return DB.Model(entity).Updates(values).Error
}

// 功能：开启一个事务（返回事务tx transaction）
// 后续操作都要基于这个事务tx而不是DB！（不能直接使用封装的方法！）
// 提交事务 tx.Commit()
// 回滚事务 tx.Rollback()
func Begin() (tx *gorm.DB) {
	return DB.Begin()
}
