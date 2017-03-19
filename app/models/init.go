package models

import "os"
import "fmt"
import "strings"
import _ "github.com/go-sql-driver/mysql"
import "github.com/go-xorm/xorm"
import "github.com/revel/revel"
import "github.com/revel/config"

//数据库链接

var DB_Read *xorm.Engine //读数据

var DB_Write *xorm.Engine //写数据

func init() {
	revel.OnAppStart(InitDB)
}

//设置数据库

func InitDB() {
	//判断是否是系统的分隔符
	separator := "/"
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {
		separator = "/"
	}

	config_file := (revel.BasePath + "/conf/databases.conf")
	config_file = strings.Replace(config_file, "/", separator, -1)
	c, _ := config.ReadDefault(config_file)

	read_driver, _ := c.String("databases", "db.read.driver")
	read_dbname, _ := c.String("databases", "db.read.dbname")
	read_user, _ := c.String("databases", "db.read.user")
	read_password, _ := c.String("databases", "db.read.password")
	read_host, _ := c.String("databases", "db.read.host")

	//数据库连接
	var err error
	DB_Read, err = xorm.NewEngine(read_driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", read_user, read_password, read_host, read_dbname))
	if err != nil {
		revel.WARN.Println("DB_Read错误: %v", err)
	}
	DB_Read.ShowSQL(true)
	write_driver, _ := c.String("databases", "db.write.driver")
	write_dbname, _ := c.String("databases", "db.write.dbname")
	write_user, _ := c.String("databases", "db.write.user")
	write_password, _ := c.String("databases", "db.write.password")
	write_host, _ := c.String("databases", "db.write.host")

	DB_Write, err = xorm.NewEngine(write_driver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", write_user, write_password, write_host, write_dbname))
	if err != nil {
		revel.WARN.Println("DB_Write错误: %v", err)
	}
	DB_Write.ShowSQL(true)
}
