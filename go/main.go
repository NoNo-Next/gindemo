package main

import (
	"fmt"
	"gindemo/go/dao"
	"gindemo/go/entity"
	"gindemo/go/inits"
	"gindemo/go/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 数据库加载
	dberr := dao.InitMySql()
	if dberr != nil {
		return
	}
	defer dao.Close()
	//绑定模型
	dao.SqlSession.AutoMigrate(&entity.User{})

	//初始化casbin
	inits.InitCasbin()
	//logrus 初始化
	inits.InitLocalLogger()

	r := gin.Default()
	// gin + logrus 日志输出
	r.Use(inits.LoggerToFile())

	//登录认证中间件
	//r.Use(inits.JWTAuthMiddleware())

	//权限认证中间件
	//r.Use(inits.AuthCasbin())
	//r.StaticFile("/favicon.ico", "./favicon.ico")
	r.Static("/static", "./static")

	//r.StaticFS("/favicon.ico" , http.Dir("./favicon.ico"))
	//r.LoadHTMLGlob("tem/*")
	r.LoadHTMLGlob("tem/**/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.html", gin.H{"name": "blue"})
	})

	/*r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/user.html", nil)
	})*/

	// 添加路由
	routers.UserController(r)

	// 注册路由
	// 生成token的请求
	r.POST("/auth", inits.AuthHandler)
	// home路由需要注册认证中间件.这里使用局部中间件
	r.GET("/home", inits.JWTAuthMiddleware(), inits.HomeHandler)

	err := r.Run(":8081")
	if err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
		return
	}
}
