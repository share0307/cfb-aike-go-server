package dao

import (
	"aike-cfb-server/kernel/base"
	"aike-cfb-server/module/aike/model"
	"fmt"
	"math/rand"
)

type UserDao struct {
	base.BaseDao
}

/**
	查找全部
 */
func (user *UserDao)Find()  {
	var users []*model.UserModel

	db.Server.Select("*").Find(&users)

	fmt.Println(users)
}

/**
	数据写入
 */
func (user *UserDao)Sore() {
	userModel := new(model.UserModel)

	userModel.NickName = fmt.Sprintf("sb_%d", rand.Int())
	userModel.Email 		  = fmt.Sprintf("sb_%d@qq.com", rand.Int())
	userModel.Mobile       = fmt.Sprintf("%d", rand.Int())

	res := db.Server.Save(userModel)

	fmt.Println("save:", res)
}