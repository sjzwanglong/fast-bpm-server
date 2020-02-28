package db

import (
	"fast-bpm/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
	"time"
)

type Factory struct {
	Pool *gorm.DB
	Cfg *Config
}

var factory = &Factory{}

type Config struct {
	Url          string        `ini:"url"`
	UserName     string        `ini:"username"`
	Password     string        `ini:"password"`
	MaxIdle      int           `ini:"maxIdle"`
	MaxOpenConn  int           `ini:"maxOpenConn"`
	MaxLifeTime  time.Duration `ini:"maxLifeTime"`
	LogMode      bool          `ini:"logMode"`
	DefaultOrder string        `ini:"defaultOrder"`
}

func InitFactory(dbType string) {
	cfg, err := ini.Load(utils.GetAbsPath("../config/config.ini"))
	if err != nil {
		fmt.Println("读取配置文件错误！ err=", err)
		return
	}
	cfg.BlockMode = false
	config := new(Config)
	err = cfg.Section(dbType).MapTo(config)
	if err != nil {
		fmt.Println("映射结构体错误！ err=", err)
		return
	}

	switch dbType {
	case "mysql":
		factory.Pool = NewMysqlPool(config)
		factory.Cfg = config
	}
}

func GetDBFactory() *gorm.DB {
	return factory.Pool
}

func GetDBConfig() *Config {
	return factory.Cfg
}