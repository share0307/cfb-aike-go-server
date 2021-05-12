package model

import "aike-cfb-server/kernel/base"

/**
	用户模型
 */
type UserModel struct {
	base.BaseModel
	// 用户昵称
	NickName	string 	`gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;"`
	// 邮箱
	Email				string 	`gorm:"Type:varchar(64);DEFAULT:'';NOT NULL;"`
	// 手机号码
	Mobile			string	`gorm:"Type:varchar(20);DEFAULT:'';NOT NULL;"`
}

/**
	指定表名
 */
func (UserModel) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "user"
}