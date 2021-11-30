package dao

import (
	"fmt"
	utils2 "gindemo/go/inits"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type conf struct {
	Url string `yaml:"url"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DbName string `yaml:"database"`
	Port string `yaml:"port"`
}

func (c *conf) getConf() *conf {
	readFile, err := ioutil.ReadFile("resources/application.yaml")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	err1 := yaml.Unmarshal(readFile,c)
	if err1 != nil {
		fmt.Println(err1.Error())
		return nil
	}
	return c
}

const DRIVER = "mysql"
var SqlSession *gorm.DB

func InitMySql() (err error) {
	var c conf
	conf := c.getConf()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Url,
		conf.Port,
		conf.DbName,
	)

	SqlSession, err = gorm.Open(DRIVER,dsn)
	// 开启日志
	SqlSession.LogMode(true)
	SqlSession.SetLogger(utils2.Logger())
	if err != nil {
		panic(err)
	}
	return SqlSession.DB().Ping()
}

func Close()  {
	SqlSession.Close()
}