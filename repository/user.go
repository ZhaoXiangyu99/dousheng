package repository

import (
	"douyin/utils"
	"fmt"
	"strconv"
)

type UserToken struct {
	UId int64 `json:"user_id"`
}

type User struct {
	Id            int64  `gorm:"column:id" json:"id"`
	Name          string `gorm:"column:name" json:"name"`
	Password      string `gorm:"column:password" json:"password"`
	FollowCount   int64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"`
	IsFollow      bool   `gorm:"column:is_follow" json:"is_follow"`
	Token         string `gorm:"column:token" json:"token"`
}

var LoginUser User
var RUser User

// CreateUser 创建用户
func CreateUser(name, password string) error {
	RUser = User{
		Id:       utils.GetOnlyId(),
		Name:     name,
		Password: password,
	}
	RUser.Token = strconv.FormatInt(RUser.Id, 10)
	fmt.Printf("%+v", RUser)
	if err := db.Table("users").Create(RUser).Error; err != nil {
		print(err.Error())
		return err
	}
	return nil
}

// GetUser 查询用户
func GetUser(name, password string) error {
	var user User
	if err := db.Table("users").Where("name = ?", name).Where("password = ?", password).Find(&user).Error; err != nil {
		return err
	}
	LoginUser = user
	return nil
}

// GetLoginUser 根据id查询用户
func GetLoginUser(uid int64) error {
	if err := db.Table("users").Where("id=?", uid).Find(&LoginUser).Error; err != nil {
		return err
	}
	return nil
}
