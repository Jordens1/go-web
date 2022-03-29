package model

//  默认使用的是users表
type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// 更改默认的使用表名
func (User) TableName() string {
	return "user"
}
