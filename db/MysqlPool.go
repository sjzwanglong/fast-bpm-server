package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mysql struct {
	db *gorm.DB
}

func NewMysqlPool(config *Config) *gorm.DB {
	mysql := new(Mysql)
	mysql.initPool(config)
	return mysql.db
}

func (m *Mysql) initPool(config *Config) *Mysql {
	url := fmt.Sprintf(`%s:%s@%s`, config.UserName, config.Password, config.Url)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		fmt.Println("数据库连接失败！ err=", err)
		return nil
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(config.MaxIdle)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(config.MaxIdle)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(config.MaxLifeTime)
	// 启动日志记录
	db.LogMode(config.LogMode)

	m.db = db

	return m
}
