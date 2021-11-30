package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {

	// 使用 MySQL 数据库初始化一个 gorm 适配器
	adapter , err := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/test", true)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
		return
	}
	enforcer, err := casbin.NewEnforcer("resources/rbac_model.conf", adapter)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
		return
	}

	//添加权限
	enforcer.AddPolicy("alice","data1" , "read")
	enforcer.AddPolicy("admin","/getAllUser" , "POST")
	enforcer.AddPolicy("alice","data2" , "read")
	enforcer.AddPolicy("admin","/getAllUser" , "GET")
	enforcer.AddPolicy("alice","data2" , "read")
	//可以通过enforcer进行增删改查

	//删除
	//enforcer.RemovePolicy("alice","data1" , "read")

	enforcer.UpdatePolicy([] string{"alice","data2" , "read"} , []string {"alice","data3","read"})
	sub := "alice" // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read" // 用户对资源执行的操作。

	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		// 处理err
	}

	if ok {
		fmt.Println("权限通过")
		// 允许alice读取data1
	} else {
		fmt.Println("权限不通过")
		// 拒绝请求，抛出异常
	}

}
