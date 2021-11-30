package entity

// 数据库表明自定义，表明结构体映射那张表
func (User) TableName() string {
	return "user"
}
type User struct {
	Id int `json:"id" form:"id"`
	UserName string `json:"userName" form:"username"`
	Password string `json:"password" form:"password"`
}

