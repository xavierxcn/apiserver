package storage

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type client struct {
	*gorm.DB
}

var c *client
var cOnce sync.Once

// Client 获取数据库客户端
func Client() *client {
	cOnce.Do(func() {
		c = &client{}
		err := c.initDb()
		if err != nil {
			panic(err)
		}
	})
	return c
}

// initDb 初始化mysql数据库
func (c *client) initDb() error {
	// 判断是否配置了数据库
	if viper.GetString("mysql.host") == "" {
		return fmt.Errorf("db.host is empty")
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
	)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return errors.Wrap(err, "open mysql error")
	}

	c.DB = db

	return nil
}
